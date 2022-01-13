package txfees

import (
	"container/heap"
	"errors"

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

// TxFeeQueue is the struct which will store the ordering of transactions
// in priority of txfee per cost
type TxFeeQueue struct {
	// txheap is the Heap of transactions prioritized by
	// transaction fee per cost;
	// it is a list of TxItems
	txheap TxHeap
	// refmap is a map which maps transaction hashes (in string form)
	// to the corresponding TxItem to allow easy deletion
	refmap map[string]*TxItem
	// utxoIDs is a map which maps utxoIDs to the transaction which contains it.
	// These utxoIDs are correspond to the utxos which would be consumed
	// by the transaction it maps to.
	utxoIDs map[string]string
	// queueSize is the size of TxFeeQueue. No additional transactions are
	// added if the total number reaches this level.
	queueSize int
}

// Init initializes the TxFeeQueue
func (tfq *TxFeeQueue) Init(queueSize int) error {
	if queueSize <= 0 {
		return errors.New("TxFeeQueue.Init: queueSize <= 0")
	}
	tfq.queueSize = queueSize
	tfq.txheap = make(TxHeap, 0, tfq.queueSize)
	tfq.refmap = make(map[string]*TxItem, tfq.queueSize)
	tfq.utxoIDs = make(map[string]string, tfq.queueSize)
	return nil
}

// SetQueueSize sets the queue size of TxFeeQueue.
// Note that this *does not* remove any elements
func (tfq *TxFeeQueue) SetQueueSize(queueSize int) error {
	if queueSize <= 0 {
		return errors.New("TxFeeQueue.SetQueueSize: queueSize <= 0")
	}
	tfq.queueSize = queueSize
	return nil
}

// QueueSize returns the size of TxFeeQueue
func (tfq *TxFeeQueue) QueueSize() int {
	return tfq.queueSize
}

// Add adds a txhash to the TxFeeQueue.
// If there are no consumedUTXO conflicts, we add to queue and exit;
// if there are consumedUTXO conflicts, we check values and replace only if
// value is greater.
// Here, value specifies the feeCostRatio of the corresponding tx.
func (tfq *TxFeeQueue) Add(txhash []byte, value *uint256.Uint256, utxoIDs [][]byte) (bool, error) {
	if tfq.IsFull() {
		return false, errors.New("TxFeeQueue.Add: queue is full")
	}
	if value == nil {
		return false, errors.New("TxFeeQueue.Add: value is nil")
	}
	utxoIDsCopy := [][]byte{}
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		utxoIDsCopy = append(utxoIDsCopy, utxoID)
	}
	conflictingTxHashes, conflict := tfq.ConflictingUTXOIDs(utxoIDsCopy)
	if conflict {
		// There are conflicts with utxo sets.
		// Find all txs with conflicts and their corresponding values.
		// If current value is larger than all of them, then we drop all
		// conflicts and add this one.
		// If value is less than any, we exit and do not add tx.
		maxValue := uint256.Zero()
		for k := 0; k < len(conflictingTxHashes); k++ {
			txhash := utils.CopySlice(conflictingTxHashes[k])
			item, ok := tfq.refmap[string(txhash)]
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
		tfq.bulkDrop(conflictingTxHashes)
	}
	txString := string(txhash)
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		tfq.utxoIDs[string(utxoID)] = txString
	}
	item := &TxItem{
		txhash:  utils.CopySlice(txhash),
		value:   value.Clone(),
		utxoIDs: utxoIDsCopy,
	}
	tfq.refmap[txString] = item
	heap.Push(&tfq.txheap, item)
	return true, nil
}

// ConflictingUTXOIDs determines if the proposed additional utxos conflict
// with the present set of utxos in the TxFeeQueue;
// returns total list of corresponding txhashes where there are conflicts.
func (tfq *TxFeeQueue) ConflictingUTXOIDs(utxoIDs [][]byte) ([][]byte, bool) {
	conflictingTxHashes := [][]byte{}
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		txhashString, ok := tfq.utxoIDs[string(utxoID)]
		if ok {
			conflictingTxHashes = append(conflictingTxHashes, []byte(txhashString))
		}
	}
	return conflictingTxHashes, len(conflictingTxHashes) > 0
}

// Pop returns the txhash of the highest valued item in the TxFeeQueue
func (tfq *TxFeeQueue) Pop() (*TxItem, error) {
	if tfq.IsEmpty() {
		return nil, errors.New("TxFeeQueue.Pop: queue is empty")
	}
	// Pop item from TxHeap
	item := heap.Pop(&tfq.txheap).(*TxItem)
	// Drop references to item in reference maps
	tfq.dropReferences(item)
	return item, nil
}

// Contains checks if tx is present in queue
func (tfq *TxFeeQueue) Contains(txhash []byte) bool {
	txString := string(txhash)
	_, ok := tfq.refmap[txString]
	return ok
}

// dropReferences drops all references to an item in the maps
func (tfq *TxFeeQueue) dropReferences(item *TxItem) {
	if item == nil {
		return
	}
	// Remove utxoIDs from utxoID map
	for k := 0; k < len(item.utxoIDs); k++ {
		utxoID := utils.CopySlice(item.utxoIDs[k])
		delete(tfq.utxoIDs, string(utxoID))
	}
	// Remove txhash from txhash map
	delete(tfq.refmap, string(item.txhash))
}

func (tfq *TxFeeQueue) bulkDrop(txhashes [][]byte) {
	for k := 0; k < len(txhashes); k++ {
		txhash := utils.CopySlice(txhashes[k])
		tfq.drop(txhash)
	}
}

// drop drops a tx from the TxFeeQueue and all associated utxoIDs
func (tfq *TxFeeQueue) drop(txhash []byte) {
	txString := string(txhash)
	item, ok := tfq.refmap[txString]
	if !ok {
		return
	}
	tfq.dropReferences(item)
	// Remove element from TxHeap
	_ = heap.Remove(&tfq.txheap, item.index)
}

// DropMined drops all txs which include the listed utxoIDs
func (tfq *TxFeeQueue) DropMined(utxoIDs [][]byte) {
	// Remove all utxoIDs and the txs which include them
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		refTxHash, ok := tfq.utxoIDs[string(utxoID)]
		if !ok {
			continue
		}
		// Drop the tx which contains it
		tfq.drop([]byte(refTxHash))
	}
}

// DropAll drops all txs from TxFeeQueue
func (tfq *TxFeeQueue) DropAll() {
	tfq.txheap = make(TxHeap, 0, tfq.queueSize)
	tfq.refmap = make(map[string]*TxItem, tfq.queueSize)
	tfq.utxoIDs = make(map[string]string, tfq.queueSize)
}

// IsFull returns true if queue is at capacity
func (tfq *TxFeeQueue) IsFull() bool {
	if len(tfq.refmap) >= tfq.queueSize {
		return true
	}
	return false
}

// IsEmpty returns true if queue is empty
func (tfq *TxFeeQueue) IsEmpty() bool {
	if len(tfq.refmap) == 0 {
		return true
	}
	return false
}
