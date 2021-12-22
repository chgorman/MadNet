package indexer

import (
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
)

/*
Need double index
First index is the fee ratio (transaction fee over bytes)
Second index is the size of the object

<prefix>|<feeSizeRatio>|<size>|<txHash>
  <>

iterate in fwd direction
*/

// NewTxFeeIndex makes a new TxFeeIndex object
func NewTxFeeIndex(p, pp prefixFunc) *TxFeeIndex {
	return &TxFeeIndex{p, pp}
}

// TxFeeIndex creates an index that allows objects to be indexed by
// txfee ratio; that is fee per byte
type TxFeeIndex struct {
	prefix    prefixFunc
	refPrefix prefixFunc
}

// TxFeeIndexKey is the key for the TxFeeIndex
type TxFeeIndexKey struct {
	key []byte
}

// MarshalBinary returns the byte slice for the key object
func (tfik *TxFeeIndexKey) MarshalBinary() []byte {
	return utils.CopySlice(tfik.key)
}

// UnmarshalBinary takes in a byte slice to set the key object
func (tfik *TxFeeIndexKey) UnmarshalBinary(data []byte) {
	tfik.key = utils.CopySlice(data)
}

// TxFeeIndexRefKey is the reference key for the TxFeeIndex
type TxFeeIndexRefKey struct {
	refkey []byte
}

// MarshalBinary returns the byte slice for the key object
func (tfirk *TxFeeIndexRefKey) MarshalBinary() []byte {
	return utils.CopySlice(tfirk.refkey)
}

// UnmarshalBinary takes in a byte slice to set the key object
func (tfirk *TxFeeIndexRefKey) UnmarshalBinary(data []byte) {
	tfirk.refkey = utils.CopySlice(data)
}

// Add adds a transaction to the TxFeeIndex
func (tfi *TxFeeIndex) Add(txn *badger.Txn, fee *uint256.Uint256, size uint32, txHash []byte, isCleanup bool) error {
	feeSizeRatioBytes, err := tfi.makeFeeSizeRatioBytes(fee, size, isCleanup)
	if err != nil {
		return err
	}
	tfiKey := tfi.makeKey(feeSizeRatioBytes, size, txHash)
	key := tfiKey.MarshalBinary()
	tfiRefKey := tfi.makeRefKey(txHash)
	refKey := tfiRefKey.MarshalBinary()
	refValue := tfi.makeRefValue(feeSizeRatioBytes, size)
	err = utils.SetValue(txn, refKey, refValue)
	if err != nil {
		return err
	}
	return utils.SetValue(txn, key, []byte{})
}

// Drop drops a transaction from the TxFeeIndex
func (tfi *TxFeeIndex) Drop(txn *badger.Txn, txHash []byte) error {
	tfiRefKey := tfi.makeRefKey(txHash)
	refKey := tfiRefKey.MarshalBinary()
	feeSizeRatioSizeBytes, err := utils.GetValue(txn, refKey)
	if err != nil {
		return err
	}
	feeSizeRatioBytes := feeSizeRatioSizeBytes[:4]
	sizeBytes := feeSizeRatioSizeBytes[4:]
	size, err := utils.UnmarshalUint32(sizeBytes)
	if err != nil {
		return err
	}
	err = utils.DeleteValue(txn, refKey)
	if err != nil {
		return err
	}
	tfiKey := tfi.makeKey(feeSizeRatioBytes, size, txHash)
	key := tfiKey.MarshalBinary()
	return utils.DeleteValue(txn, key)
}

// makeFeeSizeRatioBytes returns the byte slice of the feeSizeRatio.
// feeSizeRatio is a float32 object. In the case of a cleanup transaction,
// the largest possible value is used.
func (tfi *TxFeeIndex) makeFeeSizeRatioBytes(fee *uint256.Uint256, size uint32, isCleanup bool) ([]byte, error) {
	if fee == nil {
		return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeSizeRatioBytes: fee is not initialized")
	}
	if size == 0 {
		return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeSizeRatioBytes: invalid size; size is zero")
	}
	if isCleanup {
		// Cleanup transactions are special transactions with no fee.
		// These must be prioritized above all other transactions.
		if !fee.IsZero() {
			return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeSizeRatioBytes: invalid fee; cleanup transaction fee must be zero")
		}
		feeRatioBytes := make([]byte, 4)
		for k := 0; k < len(feeRatioBytes); k++ {
			feeRatioBytes[k] = 255
		}
		return feeRatioBytes, nil
	}
	feeFloat64, err := fee.ToFloat64()
	if err != nil {
		return nil, err
	}
	sizeFloat64 := float64(size)
	feeSizeRatio := feeFloat64 / sizeFloat64
	// Ensure the ratio is not too large
	if feeSizeRatio > float64(constants.MaxFeeSizeRatio) {
		feeSizeRatio = float64(constants.MaxFeeSizeRatio)
	}
	feeSizeRatioFloat32 := float32(feeSizeRatio)
	return utils.MarshalFloat32(feeSizeRatioFloat32), nil
}

func (tfi *TxFeeIndex) makeKey(feeSizeRatioBytes []byte, size uint32, txHash []byte) *TxFeeIndexKey {
	key := []byte{}
	key = append(key, tfi.prefix()...)
	key = append(key, feeSizeRatioBytes...)
	sizeBytes := utils.MarshalUint32(size)
	key = append(key, sizeBytes...)
	key = append(key, txHash...)
	tfiKey := &TxFeeIndexKey{}
	tfiKey.UnmarshalBinary(key)
	return tfiKey
}

func (tfi *TxFeeIndex) makeRefKey(txHash []byte) *TxFeeIndexRefKey {
	key := []byte{}
	key = append(key, tfi.refPrefix()...)
	key = append(key, txHash...)
	tfiRefKey := &TxFeeIndexRefKey{}
	tfiRefKey.UnmarshalBinary(key)
	return tfiRefKey
}

func (tfi *TxFeeIndex) makeRefValue(feeSizeRatioBytes []byte, size uint32) []byte {
	sizeBytes := utils.MarshalUint32(size)
	refValue := []byte{}
	refValue = append(refValue, utils.CopySlice(feeSizeRatioBytes)...)
	refValue = append(refValue, sizeBytes...)
	return refValue
}
