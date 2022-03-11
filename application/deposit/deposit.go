package deposit

import (
	"math/big"

	"github.com/MadBase/MadNet/constants/dbprefix"
	"github.com/MadBase/MadNet/errorz"

	"github.com/MadBase/MadNet/application/db"
	"github.com/MadBase/MadNet/application/indexer"
	"github.com/MadBase/MadNet/application/objs"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/logging"
	"github.com/MadBase/MadNet/utils"
	"github.com/dgraph-io/badger/v2"
	"github.com/sirupsen/logrus"
)

// Handler creates a value owner index of all deposits and allows
// these deposits to be returned for use in a transaction or for verification.
type Handler struct {
	valueIndex *indexer.ValueIndex
	//erctokenIndex *indexer.ERCTokenIndex
	IsSpent func(txn *badger.Txn, utxoID []byte) (bool, error)
	logger  *logrus.Logger
}

// Init initializes the deposit handler
func (dp *Handler) Init() {
	vidx := indexer.NewValueIndex(
		dbprefix.PrefixDepositValueKey,
		dbprefix.PrefixDepositValueRefKey,
	)
	dp.valueIndex = vidx
	dp.logger = logging.GetLogger(constants.LoggerApp)
}

// IsValid determines if the deposits in txvec are valid
func (dp *Handler) IsValid(txn *badger.Txn, txs objs.TxVec) ([]*objs.TXOut, error) {
	utxoIDs, err := txs.ConsumedUTXOIDOnlyDeposits()
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return nil, err
	}
	found, missing, spent, err := dp.Get(txn, utxoIDs)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return nil, err
	}
	if len(missing) > 0 {
		return nil, errorz.ErrInvalid{}.New("depositHandler.IsValid; a deposit is missing")
	}
	if len(spent) > 0 {
		return nil, errorz.ErrInvalid{}.New("depositHandler.IsValid; a deposit is already spent")
	}
	return found, nil
}

// Add will add a deposit to the handler
func (dp *Handler) Add(txn *badger.Txn, chainID uint32, utxoID []byte, biValue *big.Int, owner *objs.Owner) error {
	if chainID == 0 {
		return errorz.ErrInvalid{}.New("depositHandler.Add; chainID is zero")
	}
	value, err := new(uint256.Uint256).FromBigInt(biValue)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	if value.IsZero() {
		return errorz.ErrInvalid{}.New("depositHandler.Add; deposit value is zero")
	}
	utxoID = utils.CopySlice(utxoID)
	utxoID = utils.ForceSliceToLength(utxoID, constants.HashLen)
	spent, err := dp.IsSpent(txn, utxoID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	if spent {
		return errorz.ErrInvalid{}.New("depositHandler.Add; a deposit is already spent")
	}
	n2 := utils.CopySlice(utxoID)
	vso := &objs.ValueStoreOwner{}
	err = vso.NewFromOwner(owner)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	vs := &objs.ValueStore{
		VSPreImage: &objs.VSPreImage{
			TXOutIdx: constants.MaxUint32,
			Value:    value,
			ChainID:  chainID,
			Owner:    vso,
			Fee:      uint256.Zero(),
		},
		TxHash: n2,
	}
	utxo := &objs.TXOut{}
	err = utxo.NewValueStore(vs)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	err = dp.valueIndex.Add(txn, utxoID, owner, value)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	key := dp.makeKey(utxoID)
	_, err = utils.GetValue(txn, key)
	if err != nil {
		if err != badger.ErrKeyNotFound {
			utils.DebugTrace(dp.logger, err)
			return err
		}
	} else {
		return errorz.ErrInvalid{}.New("depositHandler.Add; stale")
	}
	if err := db.SetUTXO(txn, key, utxo); err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	return nil
}

/*
// AddERCToken will add an ERCToken deposit to the handler
func (dp *Handler) AddERCToken(txn *badger.Txn, chainID, exitChainID uint32, utxoID []byte, biValue *big.Int, owner *objs.Owner, scAddr []byte, biTokenID *big.Int) error {
	if chainID == 0 {
		return errorz.ErrInvalid{}.New("depositHandler.AddERCToken; chainID is zero")
	}
	if exitChainID == 0 {
		return errorz.ErrInvalid{}.New("depositHandler.AddERCToken; exitChainID is zero")
	}
	value, err := new(uint256.Uint256).FromBigInt(biValue)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	if value.IsZero() {
		return errorz.ErrInvalid{}.New("depositHandler.AddERCToken; deposit value is zero")
	}
	tokenID, err := new(uint256.Uint256).FromBigInt(biTokenID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	sca := &objs.SmartContract{}
	err = sca.UnmarshalBinary(scAddr)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	// TODO: think more about how utxoID will be computed
	//		 in order to determine if this needs to be changed.
	utxoID = utils.CopySlice(utxoID)
	utxoID = utils.ForceSliceToLength(utxoID, constants.HashLen)
	spent, err := dp.IsSpent(txn, utxoID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	if spent {
		return errorz.ErrInvalid{}.New("depositHandler.AddERCToken; a deposit is already spent")
	}
	n2 := utils.CopySlice(utxoID)
	eto := &objs.ERCTokenOwner{}
	err = eto.NewFromOwner(owner)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	erct := &objs.ERCToken{
		ERCTPreImage: &objs.ERCTPreImage{
			TXOutIdx:             constants.MaxUint32,
			Value:                value,
			ChainID:              chainID,
			ExitChainID:          exitChainID,
			Owner:                eto,
			SmartContractAddress: sca,
			TokenID:              tokenID,
			Withdraw:             false,
			Fee:                  uint256.Zero(),
		},
		TxHash: n2,
	}
	utxo := &objs.TXOut{}
	err = utxo.NewERCToken(erct)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	err = dp.erctokenIndex.Add(txn, utxoID, owner, value, sca, tokenID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	key := dp.makeKey(utxoID)
	_, err = utils.GetValue(txn, key)
	if err != nil {
		if err != badger.ErrKeyNotFound {
			utils.DebugTrace(dp.logger, err)
			return err
		}
	} else {
		return errorz.ErrInvalid{}.New("depositHandler.AddERCToken; stale")
	}
	if err := db.SetUTXO(txn, key, utxo); err != nil {
		utils.DebugTrace(dp.logger, err)
		return err
	}
	return nil
}
*/

// Remove will delete all references to a deposit from the Handler
func (dp *Handler) Remove(txn *badger.Txn, utxoID []byte) error {
	err := dp.valueIndex.Drop(txn, utxoID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
	}
	return nil
}

// GetValueForOwner allows a list of utxoIDs to be returned that are equal or
// greater than the value passed as minValue, and are owned by owner.
func (dp *Handler) GetValueForOwner(txn *badger.Txn, owner *objs.Owner, minValue *uint256.Uint256, maxCount int, lastKey []byte) ([][]byte, *uint256.Uint256, []byte, error) {
	excludeSpent := func(utxoID []byte) (bool, error) {
		return dp.IsSpent(txn, utils.CopySlice(utxoID))
	}
	return dp.valueIndex.GetValueForOwner(txn, owner, minValue, excludeSpent, maxCount, lastKey)
}

// Get returns four values <found>, <missing>, <spent>, <error>
// Found returns those deposits that are both known and unspent.
// Missing returns the utxoIDs of the missing deposits.
// Spent returns those deposits found that have been spent.
// An error indicates a failure of the logic and should halt the main.
func (dp *Handler) Get(txn *badger.Txn, utxoIDs [][]byte) ([]*objs.TXOut, [][]byte, []*objs.TXOut, error) {
	f := []*objs.TXOut{}
	m := [][]byte{}
	s := []*objs.TXOut{}
	for _, utxoID := range utxoIDs {
		found, missing, spent, err := dp.getInternal(txn, utils.CopySlice(utxoID))
		if err != nil {
			utils.DebugTrace(dp.logger, err)
			return nil, nil, nil, err
		}
		switch {
		case len(missing) > 0:
			m = append(m, utils.CopySlice(utxoID))
		case spent != nil:
			s = append(s, spent)
		case found != nil:
			f = append(f, found)
		}
	}
	return f, m, s, nil
}

// getInternal returns four values <found>, <missing>, <spent>, <error>;
// look at Get for more information.
func (dp *Handler) getInternal(txn *badger.Txn, utxoID []byte) (*objs.TXOut, []byte, *objs.TXOut, error) {
	utxoID = utils.CopySlice(utxoID)
	utxoID = utils.ForceSliceToLength(utxoID, constants.HashLen)
	n2 := utils.CopySlice(utxoID)
	key := dp.makeKey(n2)
	utxo, err := db.GetUTXO(txn, key)
	if err != nil {
		if err != badger.ErrKeyNotFound {
			utils.DebugTrace(dp.logger, err)
			return nil, nil, nil, err
		}
		return nil, utxoID, nil, nil
	}
	spent, err := dp.IsSpent(txn, utxoID)
	if err != nil {
		utils.DebugTrace(dp.logger, err)
		return nil, nil, nil, err
	}
	if spent {
		return nil, nil, utxo, nil
	}
	return utxo, nil, nil, nil
}

func (dp *Handler) makeKey(utxoID []byte) []byte {
	key := dbprefix.PrefixDeposit()
	key = append(key, utils.CopySlice(utxoID)...)
	return key
}
