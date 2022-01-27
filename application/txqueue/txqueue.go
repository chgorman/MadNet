package txqueue

import (
	"errors"

	"github.com/aalpar/deheap"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/utils"
)

/*
Here we describe an overview of the desired system.

We want to have transaction selection algorithm which
attempts to populate a block by prioritizing fee-dense transactions.
To enable this to be performed quickly, we want to have a prepared
list of transactions we can choose from.
All these transactions should be high value, easily added in order
of decreasing value, and are independent (that is, all possible subsets
will form a valid block).

Ordering the transactions in order of decreasing fee density
may be performed by using a heap.
One challenge is that this must be updated each time a new blockheader
is received; thus, after receiving a new block,
we must be able to *quickly* go through and purge the heap of all transactions
which contain consumed utxos.
In order to do this, we need to keep track of where each transaction
is currently in the queue;
thus, we also have two maps:
 *	one which maps a transaction hash to its heap value;
 *	another which maps utxo to transaction which consumes it.

We also allow the ability to clear the entire queue.
Presently, the thought is that this will be performed after a regular
block proposal; it may be determined that there are additional times
when this may be done.
One challenge is that we want to ensure that the queue is not cleared
too close to when a standard proposal will be performed.
We want to ensure that the queue is essentially full even after dropping
transactions from the blockheader just received.
*/

// TxQueue is the struct which will store the ordering of transactions
// in priority of txfee per cost
type TxQueue struct {
	// txheap is the Heap of transactions prioritized by
	// transaction fee per cost;
	// it is a list of TxItems
	txheap TxHeap
	// refmap is a map which maps transaction hashes (in string form)
	// to the corresponding Item to allow easy deletion
	refmap map[string]*Item
	// utxoIDs is a map which maps utxoIDs to the transaction which contains it.
	// These utxoIDs are correspond to the utxos which would be consumed
	// by the transaction it maps to.
	utxoIDs map[string]string
	// queueSize is the size of TxQueue. No additional transactions are
	// added if the total number reaches this level.
	queueSize int
	// queueThreshold is the size where we begin to be more selective the txs
	// we add to the queue because we are getting full
	queueThreshold int
	// queueThresholdFrac is the fraction which specifies how full the queue
	// must be before we are selective
	queueThresholdFrac float32
	// queueAcceptance is the proportionality constant above which we require
	// all txs to satisfy once we have reached the threshold in order to add
	// them to the queue.
	queueAcceptanceNum   int
	queueAcceptanceDenum int
	// feeCostSum is the total sum of feeCostRatios of every object inside
	// the queue
	feeCostSum *uint256.Uint256
	// numCleanupTxs counts the number of cleanup transactions in the queue
	numCleanupTxs int
}

// Init initializes the TxQueue
func (tq *TxQueue) Init(queueSize int) error {
	tq.queueThresholdFrac = 0.75
	if err := tq.SetQueueSize(queueSize); err != nil {
		return err
	}
	tq.queueAcceptanceNum = 5
	tq.queueAcceptanceDenum = 4
	tq.ClearTxQueue()
	return nil
}

// SetQueueSize sets the queue size of TxQueue.
// Note that this *does not* remove any elements
func (tq *TxQueue) SetQueueSize(queueSize int) error {
	if queueSize <= 0 {
		return errors.New("TxQueue.SetQueueSize: queueSize <= 0")
	}
	tq.queueSize = queueSize
	tq.queueThreshold = int(tq.queueThresholdFrac * float32(tq.queueSize))
	return nil
}

// QueueSize returns the size of TxQueue
func (tq *TxQueue) QueueSize() int {
	return tq.queueSize
}

// Add adds a txhash to the TxQueue.
// If there are no consumedUTXO conflicts, we add to queue and exit;
// if there are consumedUTXO conflicts, we check values and replace only if
// value is greater.
// Here, value specifies the feeCostRatio of the corresponding tx.
func (tq *TxQueue) Add(txhash []byte, value *uint256.Uint256, utxoIDs [][]byte, isCleanup bool) (bool, error) {
	if tq.IsFull() {
		return false, errors.New("TxQueue.Add: queue is full")
	}
	if value == nil {
		return false, errors.New("TxQueue.Add: value is nil")
	}
	if len(utxoIDs) == 0 {
		return false, errors.New("TxQueue.Add: len(uxtoIDs) == 0")
	}
	if !tq.IsEmpty() && tq.aboveThreshold() {
		// The queue is close to being full, so we are more selective
		// in the txs that we add.
		feeCostThreshold, err := tq.feeCostThreshold()
		if err != nil {
			return false, err
		}
		if value.Lte(feeCostThreshold) {
			return false, nil
		}
	}
	utxoIDsCopy := [][]byte{}
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		utxoIDsCopy = append(utxoIDsCopy, utxoID)
	}
	conflictingTxHashes, conflict := tq.ConflictingUTXOIDs(utxoIDsCopy)
	if conflict {
		// There are conflicts with utxo sets.
		// Find all txs with conflicts and their corresponding values.
		// If current value is larger than all of them, then we drop all
		// conflicts and add this one.
		// If value is less than any, we exit and do not add tx.
		maxValue := uint256.Zero()
		for k := 0; k < len(conflictingTxHashes); k++ {
			txhash := utils.CopySlice(conflictingTxHashes[k])
			item, ok := tq.refmap[string(txhash)]
			if !ok {
				continue
			}
			if item.value.Gt(maxValue) {
				maxValue.Set(item.value)
			}
		}
		if value.Lte(maxValue) {
			// The value (feeCostRatio) of the potential tx we want to add
			// is below the value of a conflicting tx;
			// thus, we do not add it to the queue.
			return false, nil
		}
		tq.BulkDrop(conflictingTxHashes)
	}
	txString := string(txhash)
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		tq.utxoIDs[string(utxoID)] = txString
	}
	item := &Item{
		txhash:    utils.CopySlice(txhash),
		value:     value.Clone(),
		utxoIDs:   utxoIDsCopy,
		isCleanup: isCleanup,
	}
	if item.isCleanup {
		tq.numCleanupTxs++
	} else {
		_, err := tq.feeCostSum.Add(tq.feeCostSum, item.value)
		if err != nil {
			return false, err
		}
	}
	tq.refmap[txString] = item
	deheap.Push(&tq.txheap, item)
	return true, nil
}

// feeCostThreshold computes the minimum feeCostRatio value which is required
// to add a tx to the queue once we have reached a threshold number of txs.
func (tq *TxQueue) feeCostThreshold() (*uint256.Uint256, error) {
	averageFeeCostRatio, err := tq.averageFeeCostRatio()
	if err != nil {
		return nil, err
	}
	if (tq.queueAcceptanceNum <= 0) || (tq.queueAcceptanceDenum <= 0) {
		return nil, errors.New("TxQueue.feeCostThreshold: invalid queueAcceptance values")
	}
	num, _ := new(uint256.Uint256).FromUint64(uint64(tq.queueAcceptanceNum))
	denum, _ := new(uint256.Uint256).FromUint64(uint64(tq.queueAcceptanceDenum))
	thresholdFeeCostRatio, err := new(uint256.Uint256).Mul(averageFeeCostRatio, num)
	if err != nil {
		return nil, err
	}
	_, err = thresholdFeeCostRatio.Div(thresholdFeeCostRatio, denum)
	if err != nil {
		return nil, err
	}
	return thresholdFeeCostRatio, nil
}

// ConflictingUTXOIDs determines if the proposed additional utxos conflict
// with the present set of utxos in the TxQueue;
// returns total list of corresponding txhashes where there are conflicts.
func (tq *TxQueue) ConflictingUTXOIDs(utxoIDs [][]byte) ([][]byte, bool) {
	conflictingTxHashes := [][]byte{}
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		txhashString, ok := tq.utxoIDs[string(utxoID)]
		if ok {
			conflictingTxHashes = append(conflictingTxHashes, []byte(txhashString))
		}
	}
	return conflictingTxHashes, len(conflictingTxHashes) > 0
}

// PopMax returns the txhash of the highest valued item in the TxQueue
func (tq *TxQueue) PopMax() (*Item, error) {
	if tq.IsEmpty() {
		return nil, errors.New("TxQueue.PopMax: queue is empty")
	}
	// Pop max item from TxHeap
	item := deheap.PopMax(&tq.txheap).(*Item)
	// Drop references to item in reference maps
	tq.dropReferences(item)
	if item.isCleanup {
		tq.numCleanupTxs--
	} else {
		// Subtract the value from the total feeCostSum
		_, err := tq.feeCostSum.Sub(tq.feeCostSum, item.value)
		if err != nil {
			return nil, err
		}
	}
	return item, nil
}

// PopMin returns the txhash of the lowest valued item in the TxQueue
func (tq *TxQueue) PopMin() (*Item, error) {
	if tq.IsEmpty() {
		return nil, errors.New("TxQueue.PopMin: queue is empty")
	}
	// Pop min item from TxHeap
	item := deheap.Pop(&tq.txheap).(*Item)
	// Drop references to item in reference maps
	tq.dropReferences(item)
	if item.isCleanup {
		tq.numCleanupTxs--
	} else {
		// Subtract the value from the total feeCostSum
		_, err := tq.feeCostSum.Sub(tq.feeCostSum, item.value)
		if err != nil {
			return nil, err
		}
	}
	return item, nil
}

// MinValue returns the minimum value in TxQueue
func (tq *TxQueue) MinValue() (*uint256.Uint256, error) {
	if tq.IsEmpty() {
		return nil, errors.New("TxQueue.MinValue: queue is empty")
	}
	minValue := new(uint256.Uint256)
	err := minValue.Set(tq.txheap[0].value)
	if err != nil {
		return nil, err
	}
	return minValue, nil
}

// Contains checks if tx is present in queue
func (tq *TxQueue) Contains(txhash []byte) bool {
	txString := string(txhash)
	_, ok := tq.refmap[txString]
	return ok
}

// dropReferences drops all references to an item in the maps
func (tq *TxQueue) dropReferences(item *Item) {
	if item == nil {
		return
	}
	// Remove utxoIDs from utxoID map
	for k := 0; k < len(item.utxoIDs); k++ {
		utxoID := utils.CopySlice(item.utxoIDs[k])
		delete(tq.utxoIDs, string(utxoID))
	}
	// Remove txhash from txhash map
	delete(tq.refmap, string(item.txhash))
}

// BulkDrop drops a collection of txs from the TxQueue
func (tq *TxQueue) BulkDrop(txhashes [][]byte) {
	for k := 0; k < len(txhashes); k++ {
		txhash := utils.CopySlice(txhashes[k])
		tq.Drop(txhash)
	}
}

// Drop drops a tx from the TxQueue and all associated utxoIDs
func (tq *TxQueue) Drop(txhash []byte) error {
	txString := string(txhash)
	item, ok := tq.refmap[txString]
	if !ok {
		return nil
	}
	tq.dropReferences(item)
	// Remove element from TxHeap
	_ = deheap.Remove(&tq.txheap, item.index)
	// Subtract the value from the total feeCostSum
	_, err := tq.feeCostSum.Sub(tq.feeCostSum, item.value)
	return err
}

// DeleteMined drops all txs which include the listed utxoIDs
func (tq *TxQueue) DeleteMined(utxoIDs [][]byte) {
	// Remove all utxoIDs and the txs which include them
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		refTxHash, ok := tq.utxoIDs[string(utxoID)]
		if !ok {
			continue
		}
		// Drop the tx which contains it
		tq.Drop([]byte(refTxHash))
	}
}

// ClearTxQueue clears the tx queue.
func (tq *TxQueue) ClearTxQueue() {
	tq.feeCostSum = uint256.Zero()
	tq.txheap = make(TxHeap, 0, tq.queueSize)
	tq.refmap = make(map[string]*Item, tq.queueSize)
	tq.utxoIDs = make(map[string]string, tq.queueSize)
}

// IsFull returns true if queue is at capacity
func (tq *TxQueue) IsFull() bool {
	return len(tq.refmap) >= tq.queueSize
}

// aboveThreshold returns true if queue is above threshold.
// At this point, we want to only prioritize txs above
func (tq *TxQueue) aboveThreshold() bool {
	return len(tq.refmap) >= tq.queueThreshold
}

// IsEmpty returns true if queue is empty
func (tq *TxQueue) IsEmpty() bool {
	return len(tq.refmap) == 0
}

// averageFeeCostRatio computes the average feeCostRatio of the txs in the queue
func (tq *TxQueue) averageFeeCostRatio() (*uint256.Uint256, error) {
	if tq == nil {
		return nil, errors.New("TxQueue.averageFeeCostRatio: tq not initialized")
	}
	if tq.feeCostSum == nil {
		return nil, errors.New("TxQueue.averageFeeCostRatio: feeCostSum is nil")
	}
	if tq.numCleanupTxs < 0 {
		return nil, errors.New("TxQueue.averageFeeCostRatio: negative cleanup txs; impossible")
	}
	numStdTxs := tq.txheap.Len() - tq.numCleanupTxs
	if numStdTxs < 0 {
		return nil, errors.New("TxQueue.averageFeeCostRatio: more cleanup txs than total txs in queue; impossible")
	}
	if numStdTxs == 0 {
		// All txs in queue are cleanup txs, so the average is zero
		return uint256.Zero(), nil
	}
	uintNumStdTxs, _ := new(uint256.Uint256).FromUint64(uint64(numStdTxs))
	return new(uint256.Uint256).Div(tq.feeCostSum, uintNumStdTxs)
}
