package indexer

import (
	"github.com/MadBase/MadNet/application/objs"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
)

/*
== BADGER KEYS ==

lookup:
key: <prefix>|<owner>|<scaTokenID>|<value>|<utxoID>
value: <utxoID>

reverse lookup:
key: <prefix>|<utxoID>
value: <owner>|<scaTokenID>|<value>|<utxoID>
*/

func NewERCTokenIndex(p, pp prefixFunc) *ERCTokenIndex {
	return &ERCTokenIndex{p, pp}
}

// ERCTokenIndex creates an index that allows ERCToken objects
// to be referenced by owner
type ERCTokenIndex struct {
	prefix    prefixFunc
	refPrefix prefixFunc
}

type ERCTokenIndexKey struct {
	key []byte
}

func (eik *ERCTokenIndexKey) MarshalBinary() []byte {
	return utils.CopySlice(eik.key)
}

func (eik *ERCTokenIndexKey) UnmarshalBinary(data []byte) {
	eik.key = utils.CopySlice(data)
}

type ERCTokenIndexRefKey struct {
	refkey []byte
}

// MarshalBinary returns the byte slice for the key object
func (eirk *ERCTokenIndexRefKey) MarshalBinary() []byte {
	return utils.CopySlice(eirk.refkey)
}

// UnmarshalBinary takes in a byte slice to set the key object
func (eirk *ERCTokenIndexRefKey) UnmarshalBinary(data []byte) {
	eirk.refkey = utils.CopySlice(data)
}

// Add adds an item to the list
func (ei *ERCTokenIndex) Add(txn *badger.Txn, utxoID []byte, owner *objs.Owner, value *uint256.Uint256, sca *objs.SmartContract, tokenID *uint256.Uint256) error {
	eiKey, err := ei.makeKey(owner, value, sca, tokenID, utxoID)
	if err != nil {
		return err
	}
	key := eiKey.MarshalBinary()
	viRefKey := ei.makeRefKey(utxoID)
	refKey := viRefKey.MarshalBinary()
	ercTokenIndex, err := ei.makeERCTokenIndex(owner, value, sca, tokenID, utxoID)
	if err != nil {
		return err
	}
	err = utils.SetValue(txn, refKey, ercTokenIndex)
	if err != nil {
		return err
	}
	return utils.SetValue(txn, key, utils.CopySlice(utxoID))
}

// Drop returns a list of all txHashes that should be dropped
func (ei *ERCTokenIndex) Drop(txn *badger.Txn, utxoID []byte) error {
	utxoIDCopy := utils.CopySlice(utxoID)
	viRefKey := ei.makeRefKey(utxoIDCopy)
	refKey := viRefKey.MarshalBinary()
	ercTokenIndex, err := utils.GetValue(txn, refKey)
	if err != nil {
		return err
	}
	key := []byte{}
	key = append(key, ei.prefix()...)
	key = append(key, ercTokenIndex...)
	err = utils.DeleteValue(txn, refKey)
	if err != nil {
		return err
	}
	return utils.DeleteValue(txn, key)
}

func (ei *ERCTokenIndex) GetValueForOwner(txn *badger.Txn, owner *objs.Owner, sca *objs.SmartContract, tokenID *uint256.Uint256, minValue *uint256.Uint256, excludeFn func([]byte) (bool, error), maxCount int, lastKey []byte) ([][]byte, *uint256.Uint256, []byte, error) {
	ownerBytes, err := owner.MarshalBinary()
	if err != nil {
		return nil, nil, nil, err
	}
	scaBytes, err := sca.MarshalBinary()
	if err != nil {
		return nil, nil, nil, err
	}
	tokenIDBytes, err := tokenID.MarshalBinary()
	if err != nil {
		return nil, nil, nil, err
	}

	prefix := ei.prefix()
	prefix = append(prefix, ownerBytes...)
	prefix = append(prefix, scaBytes...)
	prefix = append(prefix, tokenIDBytes...)
	opts := badger.DefaultIteratorOptions
	opts.Prefix = prefix
	iter := txn.NewIterator(opts)
	defer iter.Close()

	result := [][]byte{}
	totalValue := uint256.Zero()
	prefixLen := len(prefix)

	if lastKey != nil {
		iter.Seek(lastKey)
		if !iter.ValidForPrefix(prefix) {
			return result, totalValue, nil, nil
		}
		iter.Next()
	} else {
		iter.Seek(prefix)
	}

	for ; iter.ValidForPrefix(prefix); iter.Next() {
		itm := iter.Item()
		key := itm.KeyCopy(nil)

		valueBytes := key[prefixLen : len(key)-constants.HashLen]
		value := &uint256.Uint256{}
		err := value.UnmarshalBinary(valueBytes)
		if err != nil {
			return nil, nil, nil, err
		}

		utxoID, err := itm.ValueCopy(nil)
		if err != nil {
			return nil, nil, nil, err
		}

		if excludeFn != nil {
			shouldExclude, err := excludeFn(utxoID)
			if err != nil {
				return nil, nil, nil, err
			}
			if shouldExclude {
				continue
			}
		}

		totalValue, err = totalValue.Add(totalValue, value)
		if err != nil {
			return nil, nil, nil, err
		}

		result = append(result, utxoID)

		if totalValue.Gte(minValue) {
			break
		}
		if len(result) >= maxCount {
			return result, totalValue, key, nil
		}
	}
	return result, totalValue, nil, nil
}

func (ei *ERCTokenIndex) makeKey(owner *objs.Owner, value *uint256.Uint256, sca *objs.SmartContract, tokenID *uint256.Uint256, utxoID []byte) (*ERCTokenIndexKey, error) {
	ercTokenIndex, err := ei.makeERCTokenIndex(owner, value, sca, tokenID, utxoID)
	if err != nil {
		return nil, err
	}
	key := []byte{}
	key = append(key, ei.prefix()...)
	key = append(key, ercTokenIndex...)
	eiKey := &ERCTokenIndexKey{}
	eiKey.UnmarshalBinary(key)
	return eiKey, nil
}

func (ei *ERCTokenIndex) makeRefKey(utxoID []byte) *ERCTokenIndexRefKey {
	refKey := []byte{}
	refKey = append(refKey, ei.refPrefix()...)
	refKey = append(refKey, utils.CopySlice(utxoID)...)
	eiRefKey := &ERCTokenIndexRefKey{}
	eiRefKey.UnmarshalBinary(refKey)
	return eiRefKey
}

func (ei *ERCTokenIndex) makeERCTokenIndex(owner *objs.Owner, value *uint256.Uint256, sca *objs.SmartContract, tokenID *uint256.Uint256, utxoID []byte) ([]byte, error) {
	valueBytes, err := value.MarshalBinary()
	if err != nil {
		return nil, err
	}
	ownerBytes, err := owner.MarshalBinary()
	if err != nil {
		return nil, err
	}
	scaBytes, err := sca.MarshalBinary()
	if err != nil {
		return nil, err
	}
	tokenIDBytes, err := tokenID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	ercTokenIndex := []byte{}
	ercTokenIndex = append(ercTokenIndex, ownerBytes...)
	ercTokenIndex = append(ercTokenIndex, scaBytes...)
	ercTokenIndex = append(ercTokenIndex, tokenIDBytes...)
	ercTokenIndex = append(ercTokenIndex, valueBytes...)
	ercTokenIndex = append(ercTokenIndex, utils.CopySlice(utxoID)...)
	return ercTokenIndex, nil
}
