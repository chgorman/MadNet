package indexer

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/crypto"
	"github.com/dgraph-io/badger/v2"
)

func makeTxFeeIndex() *TxFeeIndex {
	prefix1 := func() []byte {
		return []byte("ze")
	}
	prefix2 := func() []byte {
		return []byte("zf")
	}
	index := NewTxFeeIndex(prefix1, prefix2)
	return index
}

func TestTxFeeIndexAdd(t *testing.T) {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}()
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index := makeTxFeeIndex()
	fee := uint256.One()
	cost := uint256.Two()
	txHash := crypto.Hasher([]byte("txHash"))
	isCleanup := false
	err = db.Update(func(txn *badger.Txn) error {
		err = index.Add(txn, fee, cost, txHash, isCleanup)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTxFeeIndexDrop(t *testing.T) {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}()
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	index := makeTxFeeIndex()
	fee := uint256.One()
	cost := uint256.Two()
	txHash := crypto.Hasher([]byte("txHash"))
	isCleanup := false

	err = db.Update(func(txn *badger.Txn) error {
		err := index.Drop(txn, txHash)
		if err == nil {
			t.Fatal("Should have raised error")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		err := index.Add(txn, fee, cost, txHash, isCleanup)
		if err != nil {
			t.Fatal(err)
		}
		err = index.Drop(txn, txHash)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTxMakeFeeCostRatioBytesGood0(t *testing.T) {
	index := makeTxFeeIndex()
	fee := uint256.Zero()
	cost := uint256.One()
	isCleanup := true

	feeCostTrue := make([]byte, 32)
	for k := 0; k < len(feeCostTrue); k++ {
		feeCostTrue[k] = 255
	}

	feeCostBytes, err := index.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(feeCostBytes, feeCostTrue) {
		t.Fatal("invalid feeCostBytes")
	}
}

func TestTxMakeFeeCostRatioBytesGood1(t *testing.T) {
	index := makeTxFeeIndex()
	fee, err := new(uint256.Uint256).FromUint64(1234567890)
	if err != nil {
		t.Fatal(err)
	}
	cost, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	isCleanup := false

	value := &uint256.Uint256{}
	err = value.Set(fee)
	if err != nil {
		t.Fatal(err)
	}
	_, err = value.Mul(value, uint256.TwoPower64())
	if err != nil {
		t.Fatal(err)
	}
	_, err = value.Div(value, cost)
	if err != nil {
		t.Fatal(err)
	}
	feeCostTrue, err := value.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	feeCostBytes, err := index.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(feeCostBytes, feeCostTrue) {
		t.Logf("ret:  %x\n", feeCostBytes)
		t.Logf("true: %x\n", feeCostTrue)
		t.Fatal("invalid feeCostBytes")
	}
}

func TestTxMakeFeeCostRatioBytesGood2(t *testing.T) {
	index := makeTxFeeIndex()
	fee := uint256.Max()
	cost, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	isCleanup := false

	// The fee is flushed to 2^128 because it is larger than that.
	value := uint256.TwoPower128()
	_, err = value.Mul(value, uint256.TwoPower64())
	if err != nil {
		t.Fatal(err)
	}
	_, err = value.Div(value, cost)
	if err != nil {
		t.Fatal(err)
	}
	feeCostTrue, err := value.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	feeCostBytes, err := index.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(feeCostBytes, feeCostTrue) {
		t.Logf("ret:  %x\n", feeCostBytes)
		t.Logf("true: %x\n", feeCostTrue)
		t.Fatal("invalid feeCostBytes")
	}
}

func TestTxMakeFeeCostRatioBytesBad1(t *testing.T) {
	index := makeTxFeeIndex()
	fee := new(uint256.Uint256)
	cost := new(uint256.Uint256)
	isCleanup := false

	_, err := index.makeFeeCostRatioBytes(nil, cost, isCleanup)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	_, err = index.makeFeeCostRatioBytes(fee, nil, isCleanup)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	_, err = index.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestTxMakeFeeCostRatioBytesBad2(t *testing.T) {
	index := makeTxFeeIndex()
	fee := uint256.One()
	cost := uint256.One()
	isCleanup := true

	_, err := index.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
}

func TestTxFeeIndexMakeKey(t *testing.T) {
	index := makeTxFeeIndex()
	txHash := crypto.Hasher([]byte("txHash"))
	feeRatioBytes := make([]byte, 32)
	for k := 0; k < len(feeRatioBytes); k++ {
		feeRatioBytes[k] = byte(k)
	}
	costBytes := make([]byte, 32)
	for k := 0; k < len(costBytes); k++ {
		costBytes[k] = byte(k + 128)
	}

	trueKey := []byte{}
	trueKey = append(trueKey, index.prefix()...)
	trueKey = append(trueKey, feeRatioBytes...)
	trueKey = append(trueKey, txHash...)
	tfiKey := index.makeKey(feeRatioBytes, txHash)
	key := tfiKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not match")
	}
}

func TestTxFeeIndexMakeRefKey(t *testing.T) {
	index := makeTxFeeIndex()
	txHash := crypto.Hasher([]byte("txHash"))

	trueKey := []byte{}
	trueKey = append(trueKey, index.refPrefix()...)
	trueKey = append(trueKey, txHash...)
	tfiKey := index.makeRefKey(txHash)
	key := tfiKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not match")
	}
}
