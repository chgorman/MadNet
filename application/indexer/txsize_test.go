package indexer

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
)

func makeTxSizeIndex() *TxSizeIndex {
	prefix1 := func() []byte {
		return []byte("zg")
	}
	prefix2 := func() []byte {
		return []byte("zh")
	}
	index := NewTxSizeIndex(prefix1, prefix2)
	return index
}

func TestTxSizeIndexAdd(t *testing.T) {
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

	index := makeTxSizeIndex()
	size := uint32(1024)
	txHash := crypto.Hasher([]byte("txHash"))
	err = db.Update(func(txn *badger.Txn) error {
		err = index.Add(txn, size, txHash)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTxSizeIndexDrop(t *testing.T) {
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

	index := makeTxSizeIndex()
	size := uint32(1024)
	txHash := crypto.Hasher([]byte("txHash"))

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
		err := index.Add(txn, size, txHash)
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

func TestTxSizeIndexMakeKey(t *testing.T) {
	index := makeTxSizeIndex()
	size := uint32(1024)
	txHash := crypto.Hasher([]byte("txHash"))

	trueKey := []byte{}
	trueKey = append(trueKey, index.prefix()...)
	trueKey = append(trueKey, utils.MarshalUint32(size)...)
	trueKey = append(trueKey, txHash...)
	tsiKey := index.makeKey(size, txHash)
	key := tsiKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not match")
	}
}

func TestTxSizeIndexMakeRefKey(t *testing.T) {
	index := makeTxSizeIndex()
	txHash := crypto.Hasher([]byte("txHash"))

	trueKey := []byte{}
	trueKey = append(trueKey, index.refPrefix()...)
	trueKey = append(trueKey, txHash...)
	tsiKey := index.makeRefKey(txHash)
	key := tsiKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not match")
	}
}
