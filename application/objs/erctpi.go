package objs

import (
	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/application/objs/erctpreimage"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	capnp "github.com/MadBase/go-capnproto2/v2"
)

// ERCTPreImage is a value store preimage
type ERCTPreImage struct {
	ChainID              uint32
	ExitChainID          uint32
	Value                *uint256.Uint256
	TXOutIdx             uint32
	Withdraw             bool
	Owner                *ERCTokenOwner
	Fee                  *uint256.Uint256
	TokenID              *uint256.Uint256
	SmartContractAddress *SmartContract
	//
	preHash []byte
}

// UnmarshalBinary takes a byte slice and returns the corresponding
// VSPreImage object
func (b *ERCTPreImage) UnmarshalBinary(data []byte) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erctpi.UnmarshalBinary: erctpi not initialized")
	}
	bc, err := erctpreimage.Unmarshal(data)
	if err != nil {
		return err
	}
	return b.UnmarshalCapn(bc)
}

// MarshalBinary takes the VSPreImage object and returns the canonical
// byte slice
func (b *ERCTPreImage) MarshalBinary() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erctpi.MarshalBinary: erctpi not initialized")
	}
	bc, err := b.MarshalCapn(nil)
	if err != nil {
		return nil, err
	}
	return erctpreimage.Marshal(bc)
}

// UnmarshalCapn unmarshals the capnproto definition of the object
func (b *ERCTPreImage) UnmarshalCapn(bc mdefs.ERCTPreImage) error {
	if err := erctpreimage.Validate(bc); err != nil {
		return err
	}
	b.ChainID = bc.ChainID()
	b.ExitChainID = bc.ExitChainID()
	u32array := [8]uint32{}
	u32array[0] = bc.Value0()
	u32array[1] = bc.Value1()
	u32array[2] = bc.Value2()
	u32array[3] = bc.Value3()
	u32array[4] = bc.Value4()
	u32array[5] = bc.Value5()
	u32array[6] = bc.Value6()
	u32array[7] = bc.Value7()
	vObj := &uint256.Uint256{}
	err := vObj.FromUint32Array(u32array)
	if err != nil {
		return err
	}
	b.Value = vObj
	b.TXOutIdx = bc.TXOutIdx()
	b.Withdraw = bc.Withdraw()

	sc := &SmartContract{}
	if err := sc.UnmarshalBinary(bc.SmartContractAddress()); err != nil {
		return err
	}
	b.SmartContractAddress = sc

	owner := &ERCTokenOwner{}
	if err := owner.UnmarshalBinary(bc.Owner()); err != nil {
		return err
	}
	b.Owner = owner
	fObj := &uint256.Uint256{}
	u32array[0] = bc.Fee0()
	u32array[1] = bc.Fee1()
	u32array[2] = bc.Fee2()
	u32array[3] = bc.Fee3()
	u32array[4] = bc.Fee4()
	u32array[5] = bc.Fee5()
	u32array[6] = bc.Fee6()
	u32array[7] = bc.Fee7()
	err = fObj.FromUint32Array(u32array)
	if err != nil {
		return err
	}
	b.Fee = fObj
	tidObj := &uint256.Uint256{}
	u32array[0] = bc.TokenID0()
	u32array[1] = bc.TokenID1()
	u32array[2] = bc.TokenID2()
	u32array[3] = bc.TokenID3()
	u32array[4] = bc.TokenID4()
	u32array[5] = bc.TokenID5()
	u32array[6] = bc.TokenID6()
	u32array[7] = bc.TokenID7()
	err = tidObj.FromUint32Array(u32array)
	if err != nil {
		return err
	}
	b.TokenID = tidObj
	return nil
}

// MarshalCapn marshals the object into its capnproto definition
func (b *ERCTPreImage) MarshalCapn(seg *capnp.Segment) (mdefs.ERCTPreImage, error) {
	if b == nil {
		return mdefs.ERCTPreImage{}, errorz.ErrInvalid{}.New("erctpi.MarshalCapn: erctpi not initialized")
	}
	var bc mdefs.ERCTPreImage
	if seg == nil {
		_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			return bc, err
		}
		tmp, err := mdefs.NewRootERCTPreImage(seg)
		if err != nil {
			return bc, err
		}
		bc = tmp
	} else {
		tmp, err := mdefs.NewERCTPreImage(seg)
		if err != nil {
			return bc, err
		}
		bc = tmp
	}
	owner, err := b.Owner.MarshalBinary()
	if err != nil {
		return bc, err
	}
	if err := bc.SetOwner(owner); err != nil {
		return bc, err
	}
	sca, err := b.SmartContractAddress.MarshalBinary()
	if err != nil {
		return bc, err
	}
	if err := bc.SetSmartContractAddress(sca); err != nil {
		return bc, err
	}
	bc.SetChainID(b.ChainID)
	bc.SetExitChainID(b.ExitChainID)
	u32array, err := b.Value.ToUint32Array()
	if err != nil {
		return bc, err
	}
	bc.SetValue0(u32array[0])
	bc.SetValue1(u32array[1])
	bc.SetValue2(u32array[2])
	bc.SetValue3(u32array[3])
	bc.SetValue4(u32array[4])
	bc.SetValue5(u32array[5])
	bc.SetValue6(u32array[6])
	bc.SetValue7(u32array[7])
	u32array, err = b.Fee.ToUint32Array()
	if err != nil {
		return bc, err
	}
	bc.SetFee0(u32array[0])
	bc.SetFee1(u32array[1])
	bc.SetFee2(u32array[2])
	bc.SetFee3(u32array[3])
	bc.SetFee4(u32array[4])
	bc.SetFee5(u32array[5])
	bc.SetFee6(u32array[6])
	bc.SetFee7(u32array[7])
	u32array, err = b.TokenID.ToUint32Array()
	if err != nil {
		return bc, err
	}
	bc.SetTokenID0(u32array[0])
	bc.SetTokenID1(u32array[1])
	bc.SetTokenID2(u32array[2])
	bc.SetTokenID3(u32array[3])
	bc.SetTokenID4(u32array[4])
	bc.SetTokenID5(u32array[5])
	bc.SetTokenID6(u32array[6])
	bc.SetTokenID7(u32array[7])
	bc.SetTXOutIdx(b.TXOutIdx)
	bc.SetWithdraw(b.Withdraw)
	return bc, nil
}

// PreHash calculates the PreHash of the object
func (b *ERCTPreImage) PreHash() ([]byte, error) {
	if b == nil {
		return nil, errorz.ErrInvalid{}.New("erctpi.PreHash: erctpi not initialized")
	}
	if b.preHash != nil {
		return utils.CopySlice(b.preHash), nil
	}
	msg, err := b.MarshalBinary()
	if err != nil {
		return nil, err
	}
	hsh := crypto.Hasher(msg)
	b.preHash = hsh
	return utils.CopySlice(b.preHash), nil
}

// ValidateSignature validates the signature for VSPreImage
func (b *ERCTPreImage) ValidateSignature(msg []byte, sig *ERCTokenSignature) error {
	if b == nil {
		return errorz.ErrInvalid{}.New("erctpi.ValidateSignature: erctpi not initialized")
	}
	return b.Owner.ValidateSignature(msg, sig)
}

// SmartContract is a value store preimage
type SmartContract struct {
	address []byte
}

// MarshalBinary marshals the smart contract object
func (sc *SmartContract) MarshalBinary() ([]byte, error) {
	if sc == nil {
		return nil, errorz.ErrInvalid{}.New("sc.MarshalBinary: sc not initialized")
	}
	if len(sc.address) != 20 {
		return nil, errorz.ErrInvalid{}.New("sc.MarshalBinary: sc.address has incorrect length")
	}
	return utils.CopySlice(sc.address), nil
}

// UnmarshalBinary unmarshals the smart contract object
func (sc *SmartContract) UnmarshalBinary(data []byte) error {
	if sc == nil {
		return errorz.ErrInvalid{}.New("sc.UnmarshalBinary: sc not initialized")
	}
	if len(data) != 20 {
		return errorz.ErrInvalid{}.New("sc.UnmarshalBinary: data has incorrect length")
	}
	sc.address = utils.CopySlice(data)
	return nil
}
