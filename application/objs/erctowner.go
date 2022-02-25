package objs

import (
	"bytes"

	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/errorz"
	"github.com/MadBase/MadNet/utils"
)

// ERCTokenOwner contains information related to the owner of the ERCToken
type ERCTokenOwner struct {
	SVA       SVA
	CurveSpec constants.CurveSpec
	Account   []byte
}

// New makes a new ERCTokenOwner
func (eto *ERCTokenOwner) New(acct []byte, curveSpec constants.CurveSpec) {
	eto.SVA = ERCTokenSVA
	eto.CurveSpec = curveSpec
	eto.Account = utils.CopySlice(acct)
}

// NewFromOwner takes an Owner object and creates the corresponding
// ERCTokenOwner
func (eto *ERCTokenOwner) NewFromOwner(o *Owner) error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.NewFromOwner; eto not initialized")
	}
	if err := o.Validate(); err != nil {
		return err
	}
	eto.New(o.Account, o.CurveSpec)
	if err := eto.Validate(); err != nil {
		eto.SVA = 0
		eto.CurveSpec = 0
		eto.Account = nil
		return err
	}
	return nil
}

// MarshalBinary takes the ERCTokenOwner object and returns the canonical
// byte slice
func (eto *ERCTokenOwner) MarshalBinary() ([]byte, error) {
	if err := eto.Validate(); err != nil {
		return nil, err
	}
	owner := []byte{}
	owner = append(owner, []byte{uint8(eto.SVA)}...)
	owner = append(owner, []byte{uint8(eto.CurveSpec)}...)
	owner = append(owner, utils.CopySlice(eto.Account)...)
	return owner, nil
}

// Validate validates the ERCTokenOwner object
func (eto *ERCTokenOwner) Validate() error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.Validate; eto not initialized")
	}
	if err := eto.validateSVA(); err != nil {
		return err
	}
	if err := eto.validateCurveSpec(); err != nil {
		return err
	}
	if err := eto.validateAccount(); err != nil {
		return err
	}
	return nil
}

// UnmarshalBinary takes a byte slice and returns the corresponding
// ERCTokenOwner object
func (eto *ERCTokenOwner) UnmarshalBinary(o []byte) error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.UnmarshalBinary; eto not initialized")
	}
	owner := utils.CopySlice(o)
	sva, owner, err := extractSVA(owner)
	if err != nil {
		return err
	}
	curveSpec, owner, err := extractCurveSpec(owner)
	if err != nil {
		return err
	}
	account, owner, err := extractAccount(owner)
	if err != nil {
		return err
	}
	if err := extractZero(owner); err != nil {
		return err
	}
	eto.SVA = sva
	eto.CurveSpec = curveSpec
	eto.Account = account
	if err := eto.Validate(); err != nil {
		return err
	}
	return nil
}

// ValidateSignature validates ERCTokenSignature sig for message msg
func (eto *ERCTokenOwner) ValidateSignature(msg []byte, sig *ERCTokenSignature) error {
	if err := eto.Validate(); err != nil {
		return errorz.ErrInvalid{}.New("eto.ValidateSignature; invalid ERCTokenOwner")
	}
	if err := sig.Validate(); err != nil {
		return errorz.ErrInvalid{}.New("eto.ValidateSignature; invalid ERCTokenSignature")
	}
	if eto.CurveSpec != sig.CurveSpec {
		return errorz.ErrInvalid{}.New("eto.ValidateSignature; mismatched curve spec")
	}
	signature := sig.Signature
	switch eto.CurveSpec {
	case constants.CurveSecp256k1:
		val := crypto.Secp256k1Validator{}
		pk, err := val.Validate(msg, signature)
		if err != nil {
			return err
		}
		account := crypto.GetAccount(pk)
		if !bytes.Equal(account, eto.Account) {
			return errorz.ErrInvalid{}.New("eto.ValidateSignature; invalid sig for secp256k1 account")
		}
		return nil
	case constants.CurveBN256Eth:
		val := crypto.BNValidator{}
		pk, err := val.Validate(msg, signature)
		if err != nil {
			return err
		}
		account := crypto.GetAccount(pk)
		if !bytes.Equal(account, eto.Account) {
			return errorz.ErrInvalid{}.New("eto.ValidateSignature; invalid sig for bn256 account")
		}
		return nil
	default:
		return errorz.ErrInvalid{}.New("eto.ValidateSignature; invalid curve spec")
	}
}

func (eto *ERCTokenOwner) validateCurveSpec() error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.validateCurveSpec; eto not initialized")
	}
	if !(eto.CurveSpec == constants.CurveSecp256k1) && !(eto.CurveSpec == constants.CurveBN256Eth) {
		return errorz.ErrInvalid{}.New("eto.validateCurveSpec; invalid curve spec")
	}
	return nil
}

func (eto *ERCTokenOwner) validateSVA() error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.validateSVA; eto not initialized")
	}
	if eto.SVA != ERCTokenSVA {
		return errorz.ErrInvalid{}.New("eto.validateSVA; invalid signature verification algorithm")
	}
	return nil
}

func (eto *ERCTokenOwner) validateAccount() error {
	if eto == nil {
		return errorz.ErrInvalid{}.New("eto.validateAccount; eto not initialized")
	}
	if len(eto.Account) != constants.OwnerLen {
		return errorz.ErrInvalid{}.New("eto.validateAccount; eto.account has incorrect length")
	}
	return nil
}

// Sign signs message msg with signer s
func (eto *ERCTokenOwner) Sign(msg []byte, s Signer) (*ERCTokenSignature, error) {
	sig := &ERCTokenSignature{
		SVA: ERCTokenSVA,
	}
	switch s.(type) {
	case *crypto.Secp256k1Signer:
		sig.CurveSpec = constants.CurveSecp256k1
		signature, err := s.Sign(msg)
		if err != nil {
			return nil, err
		}
		sig.Signature = signature
		return sig, nil
	case *crypto.BNSigner:
		sig.CurveSpec = constants.CurveBN256Eth
		signature, err := s.Sign(msg)
		if err != nil {
			return nil, err
		}
		sig.Signature = signature
		return sig, nil
	default:
		return nil, errorz.ErrInvalid{}.New("eto.Sign; invalid signer type")
	}
}

// ERCTokenSignature is a struct which the necessary information
// for signing an ERCToken
type ERCTokenSignature struct {
	SVA       SVA
	CurveSpec constants.CurveSpec
	Signature []byte
}

// UnmarshalBinary takes a byte slice and returns the corresponding
// ERCTokenSignature object
func (ets *ERCTokenSignature) UnmarshalBinary(signature []byte) error {
	if ets == nil {
		return errorz.ErrInvalid{}.New("ets.UnmarshalBinary; ets not initialized")
	}
	sva, signature, err := extractSVA(signature)
	if err != nil {
		return err
	}
	ets.SVA = sva
	curveSpec, signature, err := extractCurveSpec(signature)
	if err != nil {
		return err
	}
	ets.CurveSpec = curveSpec
	signature, null, err := extractSignature(signature, curveSpec)
	if err != nil {
		return err
	}
	ets.Signature = signature
	if err := extractZero(null); err != nil {
		return err
	}
	return ets.Validate()
}

// MarshalBinary takes the ERCTokenSignature object and returns the canonical
// byte slice
func (ets *ERCTokenSignature) MarshalBinary() ([]byte, error) {
	if err := ets.Validate(); err != nil {
		return nil, err
	}
	signature := []byte{}
	signature = append(signature, []byte{uint8(ets.SVA)}...)
	signature = append(signature, []byte{uint8(ets.CurveSpec)}...)
	signature = append(signature, utils.CopySlice(ets.Signature)...)
	return signature, nil
}

// Validate validates the ERCTokenSignature object
func (ets *ERCTokenSignature) Validate() error {
	if ets == nil {
		return errorz.ErrInvalid{}.New("ets.Validate; ets not initialized")
	}
	if err := ets.validateSVA(); err != nil {
		return err
	}
	if err := ets.validateCurveSpec(); err != nil {
		return err
	}
	return validateSignatureLen(ets.Signature, ets.CurveSpec)
}

func (ets *ERCTokenSignature) validateSVA() error {
	if ets == nil {
		return errorz.ErrInvalid{}.New("ets.validateSVA; ets not initialized")
	}
	if ets.SVA != ERCTokenSVA {
		return errorz.ErrInvalid{}.New("ets.validateSVA; invalid signature verification algorithm")
	}
	return nil
}

func (ets *ERCTokenSignature) validateCurveSpec() error {
	if ets == nil {
		return errorz.ErrInvalid{}.New("ets.validateCurveSpec; ets not initialized")
	}
	if !(ets.CurveSpec == constants.CurveSecp256k1) && !(ets.CurveSpec == constants.CurveBN256Eth) {
		return errorz.ErrInvalid{}.New("ets.validateCurveSpec; invalid curve spec")
	}
	return nil
}
