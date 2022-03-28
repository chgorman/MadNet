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

func TestERCTokenIndexGetValueForOwnerGood(t *testing.T) {
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

	minValue := uint256.Two()
	err = db.Update(func(txn *badger.Txn) error {
		utxoIDs, retValue, lastKey, err := index.GetValueForOwner(txn, owner, sc, tokenID, minValue, nil, 256, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(utxoIDs) != 0 {
			t.Fatal("Should not have returned any values")
		}
		if !retValue.IsZero() {
			t.Fatal("Returned value should be zero")
		}
		if lastKey != nil {
			t.Fatal("LastKey should be nil")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Add utxos to database
	utxoID1 := crypto.Hasher([]byte("utxoID1"))
	value1, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	utxoID2 := crypto.Hasher([]byte("utxoID2"))
	value2, err := new(uint256.Uint256).FromUint64(2)
	if err != nil {
		t.Fatal(err)
	}
	utxoID3 := crypto.Hasher([]byte("utxoID3"))
	value3, err := new(uint256.Uint256).FromUint64(3)
	if err != nil {
		t.Fatal(err)
	}
	err = db.Update(func(txn *badger.Txn) error {
		err = index.Add(txn, utxoID1, owner, value1, sc, tokenID)
		if err != nil {
			t.Fatal(err)
		}
		err = index.Add(txn, utxoID2, owner, value2, sc, tokenID)
		if err != nil {
			t.Fatal(err)
		}
		err = index.Add(txn, utxoID3, owner, value3, sc, tokenID)
		if err != nil {
			t.Fatal(err)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		utxoIDs, retValue, lastKey, err := index.GetValueForOwner(txn, owner, sc, tokenID, minValue, nil, 256, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(utxoIDs) == 0 {
			t.Fatal("Should have returned utxoIDs")
		}
		for k := 0; k < len(utxoIDs); k++ {
			if len(utxoIDs[k]) != constants.HashLen {
				t.Fatal("Invalid utxoID; incorrect length")
			}
		}
		if retValue.Lt(minValue) {
			t.Fatal("Returned value should be greater than minValue")
		}
		if lastKey != nil {
			t.Fatal("LastKey should be nil")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenIndexGetValueForOwnerBad(t *testing.T) {
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
		//_, _, _, err = index.GetValueForOwner(txn, owner, sc, tokenID, minValue, nil, 256, nil)
		_, _, _, err = index.GetValueForOwner(txn, nil, nil, nil, nil, nil, 256, nil)
		if err == nil {
			t.Fatal("Should have raised error (1)")
		}
		_, _, _, err = index.GetValueForOwner(txn, owner, nil, nil, nil, nil, 256, nil)
		if err == nil {
			t.Fatal("Should have raised error (2)")
		}
		_, _, _, err = index.GetValueForOwner(txn, owner, sc, nil, nil, nil, 256, nil)
		if err == nil {
			t.Fatal("Should have raised error (3)")
		}
		//_, _, _, err = index.GetValueForOwner(txn, owner, sc, tokenID, minValue, nil, 256, nil)
		_, _, _, err = index.GetValueForOwner(txn, owner, sc, tokenID, nil, nil, 256, nil)
		if err == nil {
			t.Fatal("Should have raised error (4)")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
