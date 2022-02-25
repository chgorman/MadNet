package objs

import (
	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/application/objs/erctoken"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/application/wrapper"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	capnp "github.com/MadBase/go-capnproto2/v2"
)

// ERCToken stores value in a UTXO
type ERCToken struct {
	ERCTPreImage *ERCTPreImage
	TxHash       []byte
	//
	utxoID []byte
}

// New creates a new ERCToken
func (b *ERCToken) New(chainID, exitChainID uint32, value, fee, tokenID *uint256.Uint256, acct []byte, curveSpec constants.CurveSpec, withdraw bool, txHash []byte, sca *SmartContract) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erct.New: vs not initialized")
	}
	if value == nil {
		return errorz.ErrInvalid{}.New("erct.New: value is nil")
	}
	if value.IsZero() {
		return errorz.ErrInvalid{}.New("erct.New: value is zero")
	}
	if fee == nil {
		return errorz.ErrInvalid{}.New("erct.New: fee is nil")
	}
	if tokenID == nil {
		return errorz.ErrInvalid{}.New("erct.New: tokenID is nil")
	}
	erctowner := &ERCTokenOwner{}
	erctowner.New(acct, curveSpec)
	if err := erctowner.Validate(); err != nil {
		return err
	}
	if chainID == 0 {
		return errorz.ErrInvalid{}.New("erct.New: ChainID is zero")
	}
	if exitChainID == 0 {
		return errorz.ErrInvalid{}.New("erct.New: ExitChainID is zero")
	}
	if len(txHash) != constants.HashLen {
		return errorz.ErrInvalid{}.New("vs.new: invalid txHash; incorrect txhash length")
	}
	erctp := &ERCTPreImage{
		ChainID:              chainID,
		ExitChainID:          exitChainID,
		TXOutIdx:             constants.MaxUint32,
		Withdraw:             withdraw,
		Value:                value.Clone(),
		SmartContractAddress: sca,
		Owner:                erctowner,
		Fee:                  fee.Clone(),
		TokenID:              tokenID.Clone(),
	}
	b.ERCTPreImage = erctp
	b.TxHash = utils.CopySlice(txHash)
	return nil
}

// NewFromDeposit creates a new ERCToken from a deposit event
func (b *ERCToken) NewFromDeposit(chainID, exitChainID uint32, value, tokenID *uint256.Uint256, acct []byte, curveSpec constants.CurveSpec, nonce []byte, sca *SmartContract) error {
	erctowner := &ERCTokenOwner{}
	erctowner.New(acct, curveSpec)
	if err := erctowner.Validate(); err != nil {
		return err
	}
	if chainID == 0 {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: ChainID is zero")
	}
	if exitChainID == 0 {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: ExitChainID is zero")
	}
	if len(nonce) != constants.HashLen {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: invalid nonce; incorrect nonce length")
	}
	if value == nil {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: value is nil")
	}
	if value.IsZero() {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: value is zero")
	}
	if tokenID == nil {
		return errorz.ErrInvalid{}.New("erct.NewFromDeposit: tokenID is nil")
	}
	erctp := &ERCTPreImage{
		ChainID:              chainID,
		ExitChainID:          exitChainID,
		Withdraw:             false,
		SmartContractAddress: sca,
		TXOutIdx:             constants.MaxUint32,
		Owner:                erctowner,
		Value:                value.Clone(),
		Fee:                  uint256.Zero(),
		TokenID:              tokenID.Clone(),
	}
	b.ERCTPreImage = erctp
	b.TxHash = utils.CopySlice(nonce)
	return nil
}

// UnmarshalBinary takes a byte slice and returns the corresponding
// ERCToken object
func (b *ERCToken) UnmarshalBinary(data []byte) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erct.UnmarshalBinary: erct not initialized")
	}
	bc, err := erctoken.Unmarshal(data)
	if err != nil {
		return err
	}
	return b.UnmarshalCapn(bc)
}

// MarshalBinary takes the ERCToken object and returns the canonical
// byte slice
func (b *ERCToken) MarshalBinary() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.MarshalBinary: erct not initialized")
	}
	bc, err := b.MarshalCapn(nil)
	if err != nil {
		return nil, err
	}
	return erctoken.Marshal(bc)
}

// UnmarshalCapn unmarshals the capnproto definition of the object
func (b *ERCToken) UnmarshalCapn(bc mdefs.ERCToken) error {
	if err := erctoken.Validate(bc); err != nil {
		return err
	}
	b.ERCTPreImage = &ERCTPreImage{}
	if err := b.ERCTPreImage.UnmarshalCapn(bc.ERCTPreImage()); err != nil {
		return err
	}
	b.TxHash = utils.CopySlice(bc.TxHash())
	return nil
}

// MarshalCapn marshals the object into its capnproto definition
func (b *ERCToken) MarshalCapn(seg *capnp.Segment) (mdefs.ERCToken, error) {
	if b == nil {
		return mdefs.ERCToken{}, errorz.ErrInvalid{}.New("erct.MarshalCapn: erct not initialized")
	}
	var bc mdefs.ERCToken
	if seg == nil {
		_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			return bc, err
		}
		tmp, err := mdefs.NewRootERCToken(seg)
		if err != nil {
			return bc, err
		}
		bc = tmp
	} else {
		tmp, err := mdefs.NewERCToken(seg)
		if err != nil {
			return bc, err
		}
		bc = tmp
	}
	seg = bc.Struct.Segment()
	bt, err := b.ERCTPreImage.MarshalCapn(seg)
	if err != nil {
		return bc, err
	}
	if err := bc.SetERCTPreImage(bt); err != nil {
		return bc, err
	}
	if err := bc.SetTxHash(utils.CopySlice(b.TxHash)); err != nil {
		return bc, err
	}
	return bc, nil
}

// PreHash calculates the PreHash of the object
func (b *ERCToken) PreHash() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("ertc.PreHash: erct not initialized")
	}
	return b.ERCTPreImage.PreHash()
}

// UTXOID calculates the UTXOID of the object
// TODO: determine if this is correct! We may need to do something special
func (b *ERCToken) UTXOID() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.utxoID: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return nil, errorz.ErrInvalid{}.New("erct.utxoID: erctpi not initialized")
	}
	if len(b.TxHash) != constants.HashLen {
		return nil, errorz.ErrInvalid{}.New("erct.utxoID: erct.txhash has incorrect length")
	}
	if b.utxoID != nil {
		return utils.CopySlice(b.utxoID), nil
	}
	b.utxoID = MakeUTXOID(b.TxHash, b.ERCTPreImage.TXOutIdx)
	return utils.CopySlice(b.utxoID), nil
}

// TxOutIdx returns the TxOutIdx of the object
func (b *ERCToken) TxOutIdx() (uint32, error) {
	if b == nil {
		return 0, errorz.ErrInvalid{}.New("erct.TxOutIdx: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return 0, errorz.ErrInvalid{}.New("erct.TxOutIdx: erctpi not initialized")
	}
	return b.ERCTPreImage.TXOutIdx, nil
}

// SetTxOutIdx sets the TxOutIdx of the object
func (b *ERCToken) SetTxOutIdx(idx uint32) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erct.SetTxOutIdx: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return errorz.ErrInvalid{}.New("erct.SetTxOutIdx: erctpi not initialized")
	}
	b.ERCTPreImage.TXOutIdx = idx
	return nil
}

// SetTxHash sets the TxHash of the object
func (b *ERCToken) SetTxHash(txHash []byte) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erct.SetTxHash: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return errorz.ErrInvalid{}.New("erct.SetTxHash: erctpi not initialized")
	}
	if len(txHash) != constants.HashLen {
		return errorz.ErrInvalid{}.New("erct.SetTxHash: invalid hash length")
	}
	b.TxHash = utils.CopySlice(txHash)
	return nil
}

// ChainID returns the ChainID of the object
func (b *ERCToken) ChainID() (uint32, error) {
	if b == nil {
		return 0, errorz.ErrInvalid{}.New("erct.ChainID: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return 0, errorz.ErrInvalid{}.New("erct.ChainID: erctpi not initialized")
	}
	if b.ERCTPreImage.ChainID == 0 {
		return 0, errorz.ErrInvalid{}.New("erct.ChainID: chainID is zero")
	}
	return b.ERCTPreImage.ChainID, nil
}

// Value returns the Value of the object
func (b *ERCToken) Value() (*uint256.Uint256, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Value: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Value: erctpi not initialized")
	}
	if b.ERCTPreImage.Value == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Value: erctpi.value not initialized")
	}
	if b.ERCTPreImage.Value.IsZero() {
		return nil, errorz.ErrInvalid{}.New("erct.Value: erctpi.value is zero")
	}
	return b.ERCTPreImage.Value.Clone(), nil
}

// Fee returns the Fee of the object
func (b *ERCToken) Fee() (*uint256.Uint256, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Fee: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Fee: erctpi not initialized")
	}
	if b.ERCTPreImage.Fee == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Fee: erctpi.fee not initialized")
	}
	return b.ERCTPreImage.Fee.Clone(), nil
}

// TokenID returns the TokenID of the object
func (b *ERCToken) TokenID() (*uint256.Uint256, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.TokenID: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return nil, errorz.ErrInvalid{}.New("erct.TokenID: erctpi not initialized")
	}
	if b.ERCTPreImage.TokenID == nil {
		return nil, errorz.ErrInvalid{}.New("erct.TokenID: erctpi.tokenID not initialized")
	}
	return b.ERCTPreImage.TokenID.Clone(), nil
}

/*
// TODO: needs fixed!
// ValuePlusFee returns the Value of the object with the associated fee
func (b *ERCToken) ValuePlusFee() (*uint256.Uint256, error) {
	value, err := b.Value()
	if err != nil {
		return nil, err
	}
	fee, err := b.Fee()
	if err != nil {
		return nil, err
	}
	total, err := new(uint256.Uint256).Add(value, fee)
	if err != nil {
		return nil, err
	}
	return total, nil
}
*/

// IsDeposit returns true if the object is a deposit
func (b *ERCToken) IsDeposit() bool {
	if b == nil || b.ERCTPreImage == nil {
		return false
	}
	return b.ERCTPreImage.TXOutIdx == constants.MaxUint32
}

// Owner returns the ERCTokenOwner of the ERCToken
func (b *ERCToken) Owner() (*ERCTokenOwner, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Owner: erct not initialized")
	}
	if b.ERCTPreImage == nil {
		return nil, errorz.ErrInvalid{}.New("erct.Owner: erctpi not initialized")
	}
	if err := b.ERCTPreImage.Owner.Validate(); err != nil {
		return nil, errorz.ErrInvalid{}.New("erct.Owner: ERCTokenOwner invalid")
	}
	return b.ERCTPreImage.Owner, nil
}

// GenericOwner returns the Owner of the ERCToken
func (b *ERCToken) GenericOwner() (*Owner, error) {
	eto, err := b.Owner()
	if err != nil {
		return nil, err
	}
	onr := &Owner{}
	if err := onr.NewFromERCTokenOwner(eto); err != nil {
		return nil, err
	}
	return onr, nil
}

// Sign generates the signature for a ERCToken at the time of consumption
func (b *ERCToken) Sign(txIn *TXIn, s Signer) error {
	if txIn == nil {
		return errorz.ErrInvalid{}.New("erct.Sign: txin not initialized")
	}
	msg, err := txIn.TXInLinker.MarshalBinary()
	if err != nil {
		return err
	}
	owner, err := b.Owner()
	if err != nil {
		return err
	}
	sig, err := owner.Sign(msg, s)
	if err != nil {
		return err
	}
	sigb, err := sig.MarshalBinary()
	if err != nil {
		return err
	}
	txIn.Signature = sigb
	return nil
}

// ValidateFee validates the fee of the object at the time of creation
func (b *ERCToken) ValidateFee(storage *wrapper.Storage) error {
	fee, err := b.Fee()
	if err != nil {
		return err
	}
	if b.IsDeposit() {
		if !fee.IsZero() {
			return errorz.ErrInvalid{}.New("erct.ValidateFee: invalid fee; deposits should have fee equal zero")
		}
		return nil
	}
	feeTrue, err := storage.GetERCTokenFee()
	if err != nil {
		return err
	}
	if fee.Cmp(feeTrue) != 0 {
		return errorz.ErrInvalid{}.New("erct.ValidateFee: invalid fee")
	}
	return nil
}

// ValidateSignature validates the signature of the ERCToken at the time of
// consumption
func (b *ERCToken) ValidateSignature(txIn *TXIn) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erct.ValidateSignature: erct not initialized")
	}
	if txIn == nil {
		return errorz.ErrInvalid{}.New("erct.ValidateSignature: txin not initialized")
	}
	msg, err := txIn.TXInLinker.MarshalBinary()
	if err != nil {
		return err
	}
	sig := &ERCTokenSignature{}
	if err := sig.UnmarshalBinary(txIn.Signature); err != nil {
		return err
	}
	return b.ERCTPreImage.ValidateSignature(msg, sig)
}

// MakeTxIn constructs a TXIn object for the current object
func (b *ERCToken) MakeTxIn() (*TXIn, error) {
	txOutIdx, err := b.TxOutIdx()
	if err != nil {
		return nil, err
	}
	cid, err := b.ChainID()
	if err != nil {
		return nil, err
	}
	if len(b.TxHash) != constants.HashLen {
		return nil, errorz.ErrInvalid{}.New("erct.MakeTxIn: invalid TxHash")
	}
	return &TXIn{
		TXInLinker: &TXInLinker{
			TXInPreImage: &TXInPreImage{
				ConsumedTxIdx:  txOutIdx,
				ConsumedTxHash: utils.CopySlice(b.TxHash),
				ChainID:        cid,
			},
		},
	}, nil
}
