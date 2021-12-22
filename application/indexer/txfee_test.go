package indexer

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/utils"
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
	size := uint32(1024)
	txHash := crypto.Hasher([]byte("txHash"))
	isCleanup := false
	err = db.Update(func(txn *badger.Txn) error {
		err = index.Add(txn, fee, size, txHash, isCleanup)
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
	size := uint32(1024)
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
		err := index.Add(txn, fee, size, txHash, isCleanup)
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

func TestTxFeeIndexMakeKey(t *testing.T) {
	index := makeTxFeeIndex()
	size := uint32(1024)
	txHash := crypto.Hasher([]byte("txHash"))
	feeRatio := float32(1) / float32(size)
	feeRatioBytes := utils.MarshalFloat32(feeRatio)

	trueKey := []byte{}
	trueKey = append(trueKey, index.prefix()...)
	trueKey = append(trueKey, feeRatioBytes...)
	trueKey = append(trueKey, utils.MarshalUint32(size)...)
	trueKey = append(trueKey, txHash...)
	tfiKey := index.makeKey(feeRatioBytes, size, txHash)
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
