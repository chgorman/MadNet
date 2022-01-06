package txfees

import (
	"container/heap"
	"errors"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/utils"
)

/*
We need the ability to store all specified txs.
This is being done with badger and will continue to be.
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
	tfq.txheap = make(TxHeap, 0, queueSize)
	tfq.refmap = make(map[string]*TxItem, queueSize)
	tfq.utxoIDs = make(map[string]string, queueSize)
	tfq.queueSize = queueSize
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

// Add adds a txhash to the TxFeeQueue
func (tfq *TxFeeQueue) Add(txhash []byte, value *uint256.Uint256, utxoIDs [][]byte) error {
	if tfq.IsFull() {
		return errors.New("TxFeeQueue.Add: queue is full")
	}
	if value == nil {
		return errors.New("TxFeeQueue.Add: value is nil")
	}
	txString := string(txhash)
	utxoIDsCopy := [][]byte{}
	// Store new utxoIDs; previous call to ValidAdd will ensure there
	// are no conflicts
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		tfq.utxoIDs[string(utxoID)] = txString
		utxoIDsCopy = append(utxoIDsCopy, utxoID)
	}
	item := &TxItem{
		txhash:  utils.CopySlice(txhash),
		value:   value.Clone(),
		utxoIDs: utxoIDsCopy,
	}
	tfq.refmap[txString] = item
	heap.Push(&tfq.txheap, item)
	return nil
}

// ValidAdd checks if any potential utxoIDs conflict with current utxoID set;
// we only tx if utxoIDs are not already present.
func (tfq *TxFeeQueue) ValidAdd(utxoIDs [][]byte) bool {
	for k := 0; k < len(utxoIDs); k++ {
		utxoID := utils.CopySlice(utxoIDs[k])
		_, ok := tfq.utxoIDs[string(utxoID)]
		if ok {
			return false
		}
	}
	return true
}

// Pop returns the highest value item from the TxFeeQueue
func (tfq *TxFeeQueue) Pop() ([]byte, error) {
	if tfq.IsEmpty() {
		return nil, errors.New("TxFeeQueue.Pop: queue is empty")
	}
	// Pop item from TxHeap
	item := heap.Pop(&tfq.txheap).(*TxItem)
	// Remove utxoIDs from utxoID map
	for k := 0; k < len(item.utxoIDs); k++ {
		utxoID := utils.CopySlice(item.utxoIDs[k])
		delete(tfq.utxoIDs, string(utxoID))
	}
	// Remove txhash from txhash map
	delete(tfq.refmap, string(item.txhash))
	return utils.CopySlice(item.txhash), nil
}

// Contains checks if tx is present in queue
func (tfq *TxFeeQueue) Contains(txhash []byte) bool {
	txString := string(txhash)
	_, ok := tfq.refmap[txString]
	return ok
}

// drop drops a tx from the TxFeeQueue and all associated utxoIDs
func (tfq *TxFeeQueue) drop(txhash []byte) {
	txString := string(txhash)
	item, ok := tfq.refmap[txString]
	if !ok {
		return
	}
	// Remove utxoIDs from utxoID map
	for k := 0; k < len(item.utxoIDs); k++ {
		utxoID := utils.CopySlice(item.utxoIDs[k])
		delete(tfq.utxoIDs, string(utxoID))
	}
	// Remove txhash from txhash map
	delete(tfq.refmap, txString)
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
		// Drop the utxoID
		delete(tfq.utxoIDs, string(utxoID))
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
