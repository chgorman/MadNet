package indexer

import (
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
)

/*
<prefix>/<size>|<txHash>
  <>

iterate in fwd direction
*/

// NewTxSizeIndex makes a new TxSizeIndex object
func NewTxSizeIndex(p, pp prefixFunc) *TxSizeIndex {
	return &TxSizeIndex{p, pp}
}

// TxSizeIndex creates an index that allows objects to be indexed by
// tx size; that is, the number of bytes when a tx is marshalled
type TxSizeIndex struct {
	prefix    prefixFunc
	refPrefix prefixFunc
}

// TxSizeIndexKey is the key for the TxSizeIndex
type TxSizeIndexKey struct {
	key []byte
}

// MarshalBinary returns the byte slice for the key object
func (tsik *TxSizeIndexKey) MarshalBinary() []byte {
	return utils.CopySlice(tsik.key)
}

// UnmarshalBinary takes in a byte slice to set the key object
func (tsik *TxSizeIndexKey) UnmarshalBinary(data []byte) {
	tsik.key = utils.CopySlice(data)
}

// TxSizeIndexRefKey is the reference key for the TxSizeIndex
type TxSizeIndexRefKey struct {
	refkey []byte
}

// MarshalBinary returns the byte slice for the key object
func (tsirk *TxSizeIndexRefKey) MarshalBinary() []byte {
	return utils.CopySlice(tsirk.refkey)
}

// UnmarshalBinary takes in a byte slice to set the key object
func (tsirk *TxSizeIndexRefKey) UnmarshalBinary(data []byte) {
	tsirk.refkey = utils.CopySlice(data)
}

// Add adds a transaction to the TxSizeIndex
func (tsi *TxSizeIndex) Add(txn *badger.Txn, size uint32, txHash []byte) error {
	tsiKey := tsi.makeKey(size, txHash)
	key := tsiKey.MarshalBinary()
	tsiRefKey := tsi.makeRefKey(txHash)
	refKey := tsiRefKey.MarshalBinary()
	refValue := tsi.makeRefValue(size)
	err := utils.SetValue(txn, refKey, refValue)
	if err != nil {
		return err
	}
	return utils.SetValue(txn, key, []byte{})
}

// Drop drops a transaction from the TxSizeIndex
func (tsi *TxSizeIndex) Drop(txn *badger.Txn, txHash []byte) error {
	tsiRefKey := tsi.makeRefKey(txHash)
	refKey := tsiRefKey.MarshalBinary()
	sizeBytes, err := utils.GetValue(txn, refKey)
	if err != nil {
		return err
	}
	size, err := utils.UnmarshalUint32(sizeBytes)
	if err != nil {
		return err
	}
	err = utils.DeleteValue(txn, refKey)
	if err != nil {
		return err
	}
	tsiKey := tsi.makeKey(size, txHash)
	key := tsiKey.MarshalBinary()
	return utils.DeleteValue(txn, key)
}

func (tsi *TxSizeIndex) makeKey(size uint32, txHash []byte) *TxSizeIndexKey {
	key := []byte{}
	key = append(key, tsi.prefix()...)
	sizeBytes := utils.MarshalUint32(size)
	key = append(key, sizeBytes...)
	key = append(key, txHash...)
	tsiKey := &TxSizeIndexKey{}
	tsiKey.UnmarshalBinary(key)
	return tsiKey
}

func (tsi *TxSizeIndex) makeRefKey(txHash []byte) *TxSizeIndexRefKey {
	key := []byte{}
	key = append(key, tsi.refPrefix()...)
	key = append(key, txHash...)
	tsiRefKey := &TxSizeIndexRefKey{}
	tsiRefKey.UnmarshalBinary(key)
	return tsiRefKey
}

func (tsi *TxSizeIndex) makeRefValue(size uint32) []byte {
	sizeBytes := utils.MarshalUint32(size)
	refValue := []byte{}
	refValue = append(refValue, sizeBytes...)
	return refValue
}

// MakeIterKey makes the prefix of the key to iterate the TxSizeIndex keyspace.
func (tsi *TxSizeIndex) MakeIterKey(size uint32) []byte {
	sizeBytes := utils.MarshalUint32(size)
	key := []byte{}
	key = append(key, tsi.prefix()...)
	key = append(key, sizeBytes...)
	return key
}
