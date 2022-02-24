package erctpreimage

import (
	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	capnp "github.com/MadBase/go-capnproto2/v2"
)

// Marshal will marshal the ERCTPreImage object.
func Marshal(v mdefs.ERCTPreImage) ([]byte, error) {
	raw, err := capnp.Canonicalize(v.Struct)
	if err != nil {
		return nil, err
	}
	out := utils.CopySlice(raw)
	return out, nil
}

// Unmarshal will unmarshal the ERCTPreImage object.
func Unmarshal(data []byte) (mdefs.ERCTPreImage, error) {
	var err error
	fn := func() (mdefs.ERCTPreImage, error) {
		defer func() {
			if r := recover(); r != nil {
				err = errorz.ErrInvalid{}.New("bad serialization")
			}
		}()
		dataCopy := utils.CopySlice(data)
		msg := &capnp.Message{Arena: capnp.SingleSegment(dataCopy)}
		obj, tmp := mdefs.ReadRootERCTPreImage(msg)
		err = tmp
		return obj, err
	}
	obj, err := fn()
	if err != nil {
		return mdefs.ERCTPreImage{}, err
	}
	return obj, nil
}

// Validate will validate the VSPreImage object
func Validate(v mdefs.ERCTPreImage) error {
	if v.ChainID() < 1 {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid; invalid ChainID")
	}
	if v.ExitChainID() < 1 {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid; invalid ExitChainID")
	}
	if !v.HasOwner() {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj does not have Owner")
	}
	if len(v.SmartContractAddress()) == 0 {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid: invalid SmartContractAddress; zero byte length")
	}
	if len(v.Owner()) == 0 {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid: invalid Owner; zero byte length")
	}
	if v.Value0()|v.Value1()|v.Value2()|v.Value3()|v.Value4()|v.Value5()|v.Value6()|v.Value7() == 0 {
		return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid; no value")
	}
	if v.TXOutIdx() != constants.MaxUint32 {
		if int(v.TXOutIdx()) >= constants.MaxTxVectorLength {
			return errorz.ErrInvalid{}.New("erctpreimage capn obj is not valid: output index is too large")
		}
	}
	return nil
}
