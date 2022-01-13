package txfees

import (
	"bytes"
	"testing"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/crypto"
)

func TestTxFeeQueueInitGood(t *testing.T) {
	queueSize := 128
	tfq := &TxFeeQueue{}
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTxFeeQueueInitBad(t *testing.T) {
	queueSize := 0
	tfq := &TxFeeQueue{}
	err := tfq.Init(queueSize)
	if err == nil {
		t.Fatal("Should have raised error")
	}
}

func TestTxFeeQueueQueueSize(t *testing.T) {
	tfq := &TxFeeQueue{}
	qs := tfq.QueueSize()
	if qs != 0 {
		t.Fatal("invalid QueueSize (1)")
	}
	queueSize := 0
	err := tfq.SetQueueSize(queueSize)
	if err == nil {
		t.Fatal("Should have raised error")
	}
	queueSize = 128
	err = tfq.SetQueueSize(queueSize)
	if err != nil {
		t.Fatal(err)
	}
	qs = tfq.QueueSize()
	if qs != queueSize {
		t.Fatal("invalid QueueSize (2)")
	}
}

func TestTxFeeQueueEmptyFull(t *testing.T) {
	tfq := &TxFeeQueue{}
	if !tfq.IsEmpty() {
		t.Fatal("Should be empty (1)")
	}
	if !tfq.IsFull() {
		t.Fatal("Should be full")
	}
	queueSize := 128
	err := tfq.SetQueueSize(queueSize)
	if err != nil {
		t.Fatal(err)
	}
	if !tfq.IsEmpty() {
		t.Fatal("Should be empty (2)")
	}
	if tfq.IsFull() {
		t.Fatal("Should not be full (1)")
	}

	err = tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID11 := crypto.Hasher([]byte("utxoID11"))
	utxoID12 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID11, utxoID12}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}

	if tfq.IsEmpty() {
		t.Fatal("Should not be empty")
	}
	if tfq.IsFull() {
		t.Fatal("Should not be full (2)")
	}
}

func TestTxFeeQueueDropAll(t *testing.T) {
	tfq := &TxFeeQueue{}
	err := tfq.Init(1)
	if err != nil {
		t.Fatal(err)
	}

	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID1"))
	utxoID2 := crypto.Hasher([]byte("utxoID2"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}

	if tfq.IsEmpty() {
		t.Fatal("Should not be empty")
	}

	tfq.DropAll()
	if len(tfq.txheap) != 0 {
		t.Fatal("TxHeap should be length 0")
	}
	if len(tfq.refmap) != 0 {
		t.Fatal("refmap should be length 0")
	}
	if len(tfq.utxoIDs) != 0 {
		t.Fatal("utxoIDs should be length 0")
	}
	if !tfq.IsEmpty() {
		t.Fatal("Should be empty")
	}
}

func TestTxFeeQueueAddBad(t *testing.T) {
	tfq := &TxFeeQueue{}
	_, err := tfq.Add(nil, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	queueSize := 128
	err = tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tfq.Add(nil, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestTxFeeQueueAddGood(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID1"))
	utxoID2 := crypto.Hasher([]byte("utxoID2"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}

	if !tfq.Contains(txhash) {
		t.Fatal("Queue should contain txhash")
	}
	if _, ok := tfq.utxoIDs[string(utxoID1)]; !ok {
		t.Fatal("utxoID1 not present")
	}
	if _, ok := tfq.utxoIDs[string(utxoID2)]; !ok {
		t.Fatal("utxoID2 not present")
	}

	tfq.drop(txhash)
	if len(tfq.txheap) != 0 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 0 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 0 {
		t.Fatal("invalid length of utxoIDs")
	}
	tfq.drop(txhash)
}

// In depth test adding items and dropping them.
// Add item1, item2, item3 (all valid additions);
// Remove item2;
// Add item2 again;
// Make item4 which contains utxoIDs from item1 and item2;
// item4 is mined, so remove all information conflicting with it.
func TestTxFeeQueueAddGood2(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add 3 txs
	txhash1 := crypto.Hasher([]byte("TxHash1"))
	utxoID11 := crypto.Hasher([]byte("utxoID11"))
	utxoID12 := crypto.Hasher([]byte("utxoID12"))
	value1, err := new(uint256.Uint256).FromUint64(19)
	if err != nil {
		t.Fatal(err)
	}
	item1 := &TxItem{
		txhash:  txhash1,
		value:   value1,
		utxoIDs: [][]byte{utxoID11, utxoID12},
	}
	if _, conflict := tfq.ConflictingUTXOIDs(item1.utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err := tfq.Add(item1.txhash, item1.value, item1.utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}

	txhash2 := crypto.Hasher([]byte("TxHash2"))
	utxoID21 := crypto.Hasher([]byte("utxoID21"))
	utxoID22 := crypto.Hasher([]byte("utxoID22"))
	value2, err := new(uint256.Uint256).FromUint64(15)
	if err != nil {
		t.Fatal(err)
	}
	item2 := &TxItem{
		txhash:  txhash2,
		value:   value2,
		utxoIDs: [][]byte{utxoID21, utxoID22},
	}
	if _, conflict := tfq.ConflictingUTXOIDs(item2.utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err = tfq.Add(item2.txhash, item2.value, item2.utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 2 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 2 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 4 {
		t.Fatal("invalid length of utxoIDs")
	}

	txhash3 := crypto.Hasher([]byte("TxHash3"))
	utxoID31 := crypto.Hasher([]byte("utxoID31"))
	utxoID32 := crypto.Hasher([]byte("utxoID32"))
	value3, err := new(uint256.Uint256).FromUint64(3)
	if err != nil {
		t.Fatal(err)
	}
	item3 := &TxItem{
		txhash:  txhash3,
		value:   value3,
		utxoIDs: [][]byte{utxoID31, utxoID32},
	}
	if _, conflict := tfq.ConflictingUTXOIDs(item3.utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err = tfq.Add(item3.txhash, item3.value, item3.utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 3 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 3 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 6 {
		t.Fatal("invalid length of utxoIDs")
	}

	// Drop item2; confirm missing
	tfq.drop(item2.txhash)
	if len(tfq.txheap) != 2 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 2 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 4 {
		t.Fatal("invalid length of utxoIDs")
	}
	if !tfq.Contains(item1.txhash) {
		t.Fatal("does not contain item1")
	}
	if tfq.Contains(item2.txhash) {
		t.Fatal("contains item2")
	}
	if !tfq.Contains(item3.txhash) {
		t.Fatal("does not contain item3")
	}

	// Add item2 again
	if _, conflict := tfq.ConflictingUTXOIDs(item2.utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err = tfq.Add(item2.txhash, item2.value, item2.utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 3 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 3 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 6 {
		t.Fatal("invalid length of utxoIDs")
	}
	if !tfq.Contains(item1.txhash) {
		t.Fatal("does not contain item1")
	}
	if !tfq.Contains(item2.txhash) {
		t.Fatal("does not contain item2")
	}
	if !tfq.Contains(item3.txhash) {
		t.Fatal("does not contain item3")
	}

	// These are part of a "new tx" which has been mined
	txhash4 := crypto.Hasher([]byte("TxHash4"))
	value4, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	newUtxoIDs := [][]byte{utxoID11, utxoID22}
	item4 := &TxItem{
		txhash:  txhash4,
		value:   value4,
		utxoIDs: newUtxoIDs,
	}
	// We should not be able to add item4 because of conflicts
	if _, conflict := tfq.ConflictingUTXOIDs(item4.utxoIDs); !conflict {
		t.Fatal("Should have conflict")
	}

	// Drop item4; confirm missing.
	// This should kick out item1 and item2.
	tfq.DropMined(item4.utxoIDs)
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}
	if tfq.Contains(item1.txhash) {
		t.Fatal("contains item1")
	}
	if tfq.Contains(item2.txhash) {
		t.Fatal("contains item2")
	}
	if !tfq.Contains(item3.txhash) {
		t.Fatal("does not contain item3")
	}
}

func TestTxFeeQueueValidAdd1(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs); !conflict {
		t.Fatal("Should have conflict")
	}
}

func TestTxFeeQueueValidAdd2(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs); conflict {
		t.Fatal("Should not have conflict")
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs); !conflict {
		t.Fatal("Should have conflict")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}
	if _, ok := tfq.refmap[string(txhash)]; !ok {
		t.Fatal("txhash1 should be present")
	}
	if _, ok := tfq.utxoIDs[string(utxoID1)]; !ok {
		t.Fatal("utxoID1 should be present")
	}
	if _, ok := tfq.utxoIDs[string(utxoID2)]; !ok {
		t.Fatal("utxoID2 should be present")
	}

	// Add another tx with higher value, kicking out the first tx
	txhash2 := crypto.Hasher([]byte("TxHash2"))
	utxoIDs2 := [][]byte{utxoID1}
	value2, err := new(uint256.Uint256).FromUint64(2)
	if err != nil {
		t.Fatal(err)
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs2); !conflict {
		t.Fatal("Should have conflict")
	}
	ok, err = tfq.Add(txhash2, value2, utxoIDs2)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 1 {
		t.Fatal("invalid length of utxoIDs")
	}
	if _, ok := tfq.refmap[string(txhash)]; ok {
		t.Fatal("txhash1 should not be present")
	}
	if _, ok := tfq.refmap[string(txhash2)]; !ok {
		t.Fatal("txhash2 should be present")
	}
	if _, ok := tfq.utxoIDs[string(utxoID1)]; !ok {
		t.Fatal("utxoID1 should be present")
	}
	if _, ok := tfq.utxoIDs[string(utxoID2)]; ok {
		t.Fatal("utxoID2 should not be present")
	}

	// Attempt to add another tx with lower value
	txhash3 := crypto.Hasher([]byte("TxHash3"))
	utxoIDs3 := [][]byte{utxoID1}
	value3, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	if _, conflict := tfq.ConflictingUTXOIDs(utxoIDs3); !conflict {
		t.Fatal("Should have conflict")
	}
	ok, err = tfq.Add(txhash3, value3, utxoIDs3)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("Should not have added")
	}
}

func TestTxFeeQueuePopBad(t *testing.T) {
	tfq := &TxFeeQueue{}
	_, err := tfq.Pop()
	if err == nil {
		t.Fatal("Should have raised error")
	}
}

func TestTxFeeQueuePopGood(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}

	retItem, err := tfq.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(retItem.txhash, txhash) {
		t.Fatal("returned hash value does not match")
	}
}

func TestTxFeeQueueDropMined1(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID1"))
	utxoIDs := [][]byte{utxoID1}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 1 {
		t.Fatal("invalid length of utxoIDs")
	}

	tfq.DropMined(utxoIDs)
	if len(tfq.txheap) != 0 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 0 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 0 {
		t.Fatal("invalid length of utxoIDs")
	}
}

func TestTxFeeQueueDropMined2(t *testing.T) {
	tfq := &TxFeeQueue{}
	queueSize := 128
	err := tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}

	// Make and add tx
	txhash := crypto.Hasher([]byte("TxHash1"))
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("did not add")
	}
	if len(tfq.txheap) != 1 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 1 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 2 {
		t.Fatal("invalid length of utxoIDs")
	}

	tfq.DropMined(utxoIDs)
	if len(tfq.txheap) != 0 {
		t.Fatal("invalid length of txheap")
	}
	if len(tfq.refmap) != 0 {
		t.Fatal("invalid length of refmap")
	}
	if len(tfq.utxoIDs) != 0 {
		t.Fatal("invalid length of utxoIDs")
	}
}

func TestTxFeeQueueDropRefernces(t *testing.T) {
	tfq := &TxFeeQueue{}
	tfq.dropReferences(nil)
}
