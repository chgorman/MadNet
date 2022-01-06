package indexer

import (
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
)

/*
Need double index
First index is the fee ratio (transaction fee per cost)
Second index is the size of the object

<prefix>|<feeCostRatio>|<size>|<txHash>
  <>

iterate in fwd direction
*/

var numBytes int = 32

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
func (tfi *TxFeeIndex) Add(txn *badger.Txn, fee *uint256.Uint256, cost *uint256.Uint256, txHash []byte, isCleanup bool) error {
	feeCostRatioBytes, err := tfi.makeFeeCostRatioBytes(fee, cost, isCleanup)
	if err != nil {
		return err
	}
	costBytes, err := cost.MarshalBinary()
	if err != nil {
		return err
	}
	tfiKey := tfi.makeKey(feeCostRatioBytes, costBytes, txHash)
	key := tfiKey.MarshalBinary()
	tfiRefKey := tfi.makeRefKey(txHash)
	refKey := tfiRefKey.MarshalBinary()
	refValue := tfi.makeRefValue(feeCostRatioBytes, costBytes)
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
	feeCostRatioSizeBytes, err := utils.GetValue(txn, refKey)
	if err != nil {
		return err
	}
	feeCostRatioBytes := feeCostRatioSizeBytes[:numBytes]
	costBytes := feeCostRatioSizeBytes[numBytes:]
	err = utils.DeleteValue(txn, refKey)
	if err != nil {
		return err
	}
	tfiKey := tfi.makeKey(feeCostRatioBytes, costBytes, txHash)
	key := tfiKey.MarshalBinary()
	return utils.DeleteValue(txn, key)
}

// makeFeeCostRatioBytes returns the byte slice of the feeCostRatio.
// feeCostRatio is a uint256 object. In the case of a cleanup transaction,
// the largest possible value is used.
func (tfi *TxFeeIndex) makeFeeCostRatioBytes(fee *uint256.Uint256, cost *uint256.Uint256, isCleanup bool) ([]byte, error) {
	if fee == nil {
		return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeCostRatioBytes: fee is not initialized")
	}
	if cost == nil {
		return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeCostRatioBytes: cost is not initialized")
	}
	if cost.IsZero() {
		return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeCostRatioBytes: cost is zero")
	}
	if isCleanup {
		// Cleanup transactions are special transactions with no fee.
		// These must be prioritized above all other transactions.
		if !fee.IsZero() {
			return nil, errorz.ErrInvalid{}.New("TxFeeIndex.makeFeeSizeRatioBytes: invalid fee; cleanup transaction fee must be zero")
		}
		feeRatioBytes := make([]byte, numBytes)
		for k := 0; k < len(feeRatioBytes); k++ {
			feeRatioBytes[k] = 255
		}
		return feeRatioBytes, nil
	}
	scaleFactor := uint256.TwoPower64()
	maxFee := uint256.TwoPower128()
	feeCopy := fee.Clone()
	// Ensure the fee is not too large
	if feeCopy.Gt(maxFee) {
		err := feeCopy.Set(maxFee)
		if err != nil {
			return nil, err
		}
	}
	scaledFee, err := new(uint256.Uint256).Mul(feeCopy, scaleFactor)
	if err != nil {
		return nil, err
	}
	feeCostRatio, err := new(uint256.Uint256).Div(scaledFee, cost)
	if err != nil {
		return nil, err
	}
	feeCostRatioBytes, err := feeCostRatio.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return feeCostRatioBytes, nil
}

func (tfi *TxFeeIndex) makeKey(feeCostRatioBytes []byte, costBytes []byte, txHash []byte) *TxFeeIndexKey {
	key := []byte{}
	key = append(key, tfi.prefix()...)
	key = append(key, feeCostRatioBytes...)
	key = append(key, costBytes...)
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

func (tfi *TxFeeIndex) makeRefValue(feeCostRatioBytes []byte, costBytes []byte) []byte {
	refValue := []byte{}
	refValue = append(refValue, utils.CopySlice(feeCostRatioBytes)...)
	refValue = append(refValue, utils.CopySlice(costBytes)...)
	return refValue
}

// NewIter makes a new iterator for iterating through TxFeeIndex
func (tfi *TxFeeIndex) NewIter(txn *badger.Txn) (*badger.Iterator, []byte) {
	prefix := tfi.prefix()
	opts := badger.DefaultIteratorOptions
	opts.Prefix = prefix
	return txn.NewIterator(opts), prefix
}
