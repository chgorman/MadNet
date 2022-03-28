package erctoken

import (
	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
	capnp "github.com/MadBase/go-capnproto2/v2"
)

// Marshal will marshal the ERCToken object.
func Marshal(v mdefs.ERCToken) ([]byte, error) {
	raw, err := capnp.Canonicalize(v.Struct)
	if err != nil {
		return nil, err
	}
	out := utils.CopySlice(raw)
	return out, nil
}

// Unmarshal will unmarshal the ERCToken object.
func Unmarshal(data []byte) (mdefs.ERCToken, error) {
	var err error
	fn := func() (mdefs.ERCToken, error) {
		defer func() {
			if r := recover(); r != nil {
				err = errorz.ErrInvalid{}.New("bad serialization")
			}
		}()
		dataCopy := utils.CopySlice(data)
		msg := &capnp.Message{Arena: capnp.SingleSegment(dataCopy)}
		obj, tmp := mdefs.ReadRootERCToken(msg)
		err = tmp
		return obj, err
	}
	obj, err := fn()
	if err != nil {
		return mdefs.ERCToken{}, err
	}
	return obj, nil
}

// Validate will validate the ERCToken object
func Validate(v mdefs.ERCToken) error {
	if !v.HasERCTPreImage() {
		return errorz.ErrInvalid{}.New("ercToken capn obj does not have ERCTPreImage")
	}
	if !v.HasTxHash() {
		return errorz.ErrInvalid{}.New("ercToken capn obj does not have TxHash")
	}
	if len(v.TxHash()) != constants.HashLen {
		return errorz.ErrInvalid{}.New("ercToken capn obj is not valid: invalid TxHash; incorrect byte length")
	}
	return nil
}
