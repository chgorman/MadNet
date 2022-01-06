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

	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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
	err := tfq.Add(nil, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	queueSize := 128
	err = tfq.Init(queueSize)
	if err != nil {
		t.Fatal(err)
	}
	err = tfq.Add(nil, nil, nil)
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
	utxoID1 := crypto.Hasher([]byte("utxoID11"))
	utxoID2 := crypto.Hasher([]byte("utxoID12"))
	utxoIDs := [][]byte{utxoID1, utxoID2}
	value, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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

func TestTxFeeQueueValidAdd(t *testing.T) {
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
	if !tfq.ValidAdd(utxoIDs) {
		t.Fatal("ValidAdd should return true")
	}
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
	}
	if tfq.ValidAdd(utxoIDs) {
		t.Fatal("ValidAdd should return false")
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
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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

	retHash, err := tfq.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(retHash, txhash) {
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
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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
	err = tfq.Add(txhash, value, utxoIDs)
	if err != nil {
		t.Fatal(err)
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
