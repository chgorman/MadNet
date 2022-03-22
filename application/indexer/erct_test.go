package indexer

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/MadBase/MadNet/application/objs"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/dgraph-io/badger/v2"
)

func makeERCTokenIndex() *ERCTokenIndex {
	prefix1 := func() []byte {
		return []byte("za")
	}
	prefix2 := func() []byte {
		return []byte("zb")
	}
	index := NewERCTokenIndex(prefix1, prefix2)
	return index
}

func TestERCTokenIndexAdd(t *testing.T) {
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

	index := makeERCTokenIndex()
	owner := &objs.Owner{}
	utxoID := crypto.Hasher([]byte("utxoID"))
	value, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	tokenID, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	sc := &objs.SmartContract{}
	addr := make([]byte, constants.OwnerLen)
	addr[0] = 127
	origChainID := uint32(128)
	err = sc.New(origChainID, addr)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := index.Add(txn, utxoID, owner, value, sc, tokenID)
		if err == nil {
			// Invalid Owner
			t.Fatal("Should have raised error (1)")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	owner = makeOwner()
	err = db.Update(func(txn *badger.Txn) error {
		err = index.Add(txn, utxoID, owner, value, sc, tokenID)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenIndexDrop(t *testing.T) {
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

	index := makeERCTokenIndex()
	owner := makeOwner()
	utxoID := crypto.Hasher([]byte("utxoID"))
	value, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	tokenID, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	sc := &objs.SmartContract{}
	addr := make([]byte, constants.OwnerLen)
	addr[0] = 127
	origChainID := uint32(128)
	err = sc.New(origChainID, addr)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := index.Drop(txn, utxoID)
		if err == nil {
			// Not present
			t.Fatal("Should have raised error")
		}
		err = index.Add(txn, utxoID, owner, value, sc, tokenID)
		if err != nil {
			t.Fatal(err)
		}
		err = index.Drop(txn, utxoID)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenIndexMakeKey(t *testing.T) {
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

	index := makeERCTokenIndex()
	owner := &objs.Owner{}
	utxoID := crypto.Hasher([]byte("utxoID"))
	value, err := new(uint256.Uint256).FromUint64(25519)
	if err != nil {
		t.Fatal(err)
	}
	tokenID, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	sc := &objs.SmartContract{}
	addr := make([]byte, constants.OwnerLen)
	addr[0] = 127
	origChainID := uint32(128)
	err = sc.New(origChainID, addr)
	if err != nil {
		t.Fatal(err)
	}

	_, err = index.makeKey(owner, value, sc, tokenID, utxoID)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	owner = makeOwner()
	ownerBytes, err := owner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	trueKey := []byte{}
	trueKey = append(trueKey, index.prefix()...)
	trueKey = append(trueKey, ownerBytes...)
	scBytes, err := sc.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	trueKey = append(trueKey, scBytes...)
	tokenIDBytes, err := tokenID.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	trueKey = append(trueKey, tokenIDBytes...)
	valueBytes, err := value.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	trueKey = append(trueKey, valueBytes...)
	trueKey = append(trueKey, utxoID...)
	viKey, err := index.makeKey(owner, value, sc, tokenID, utxoID)
	if err != nil {
		t.Fatal(err)
	}
	key := viKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not agree")
	}
}

func TestERCTokenIndexMakeRefKey(t *testing.T) {
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

	index := makeERCTokenIndex()
	utxoID := crypto.Hasher([]byte("utxoID"))
	trueKey := []byte{}
	trueKey = append(trueKey, index.refPrefix()...)
	trueKey = append(trueKey, utxoID...)
	viKey := index.makeRefKey(utxoID)
	key := viKey.MarshalBinary()
	if !bytes.Equal(key, trueKey) {
		t.Fatal("keys do not agree")
	}
}
