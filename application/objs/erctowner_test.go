package objs

import (
	"testing"

	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
)

func TestERCTokenOwnerMarshalBinary(t *testing.T) {
	erctp := &ERCTPreImage{}
	_, err := erctp.Owner.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	eto := &ERCTokenOwner{}
	_, err = eto.MarshalBinary()
	if err == nil {
		t.Fatal("Should raise an error (1)")
	}

	eto.SVA = ERCTokenSVA
	_, err = eto.MarshalBinary()
	if err == nil {
		t.Fatal("Should raise an error (2)")
	}

	eto.CurveSpec = constants.CurveSecp256k1
	_, err = eto.MarshalBinary()
	if err == nil {
		t.Fatal("Should raise an error (3)")
	}

	eto.Account = make([]byte, constants.OwnerLen)
	_, err = eto.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenOwnerUnmarshalBinary(t *testing.T) {
	data := make([]byte, 0)
	erctp := &ERCTPreImage{}
	err := erctp.Owner.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	eto := &ERCTokenOwner{}
	err = eto.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should raise an error (1)")
	}

	data = make([]byte, 1)
	data[0] = uint8(ERCTokenSVA)
	err = eto.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should raise an error (2)")
	}

	data = make([]byte, 2)
	data[0] = uint8(ERCTokenSVA)
	data[1] = uint8(constants.CurveSecp256k1)
	err = eto.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should raise an error (3)")
	}

	data = make([]byte, 23)
	data[0] = uint8(ERCTokenSVA)
	data[1] = uint8(constants.CurveSecp256k1)
	err = eto.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should raise an error (4)")
	}

	data = make([]byte, 2+constants.OwnerLen)
	data[1] = uint8(constants.CurveSecp256k1)
	err = eto.UnmarshalBinary(data)
	if err == nil {
		t.Fatal("Should raise an error (5)")
	}

	data = make([]byte, 2+constants.OwnerLen)
	data[0] = uint8(ERCTokenSVA)
	data[1] = uint8(constants.CurveSecp256k1)
	err = eto.UnmarshalBinary(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenOwnerNewFromOwner(t *testing.T) {
	o := &Owner{}
	erctp := &ERCTPreImage{}
	err := erctp.Owner.NewFromOwner(o)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	eto := &ERCTokenOwner{}
	err = eto.NewFromOwner(o)
	if err == nil {
		t.Fatal("Should raise an error (1)")
	}

	curveSpec := constants.CurveSpec(255) // bad curveSpec
	acct := make([]byte, constants.OwnerLen)
	err = o.New(acct, curveSpec)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.NewFromOwner(o)
	if err == nil {
		t.Fatal("Should raise an error (2)")
	}

	curveSpec = constants.CurveSecp256k1
	err = o.New(acct, curveSpec)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.NewFromOwner(o)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenOwnerValidate(t *testing.T) {
	erctp := &ERCTPreImage{}
	err := erctp.Owner.Validate()
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}
	err = erctp.Owner.validateSVA()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	err = erctp.Owner.validateAccount()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	err = erctp.Owner.validateCurveSpec()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	eto := &ERCTokenOwner{}
	err = eto.Validate()
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	eto.SVA = ERCTokenSVA
	err = eto.Validate()
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	eto.CurveSpec = constants.CurveSecp256k1
	err = eto.Validate()
	if err == nil {
		t.Fatal("Should have raised error (6)")
	}

	eto.Account = make([]byte, constants.OwnerLen)
	err = eto.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenOwnerSign(t *testing.T) {
	eto := &ERCTokenOwner{}
	msg := make([]byte, 0)
	signerSecp := &crypto.Secp256k1Signer{}
	_, err := eto.Sign(msg, signerSecp)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	signerBN := &crypto.BNSigner{}
	_, err = eto.Sign(msg, signerBN)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenOwnerValidateSignatureSecp(t *testing.T) {
	msg := make([]byte, 32)
	ets := &ERCTokenSignature{}
	erctp := &ERCTPreImage{}
	err := erctp.Owner.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	// Raises error for invalid eto
	eto := &ERCTokenOwner{}
	err = eto.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	// Raises error for invalid sig
	acct := make([]byte, constants.OwnerLen)
	eto.New(acct, constants.CurveSecp256k1)
	err = eto.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	// Raises error for mismatched curve spec
	secpSigner := &crypto.Secp256k1Signer{}
	privk := make([]byte, 32)
	privk[0] = 1
	privk[31] = 1
	if err := secpSigner.SetPrivk(privk); err != nil {
		t.Fatal(err)
	}
	bnSigner := &crypto.BNSigner{}
	err = bnSigner.SetPrivk(privk)
	if err != nil {
		t.Fatal(err)
	}
	etsBN, err := eto.Sign(msg, bnSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.ValidateSignature(msg, etsBN)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	msgBad := make([]byte, len(msg)+1)
	etsBad, err := eto.Sign(msgBad, secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.ValidateSignature(msg, etsBad)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	// Correct eto
	pk, err := secpSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	eto.Account = crypto.GetAccount(pk)
	ets, err = eto.Sign(msg, secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.ValidateSignature(msg, ets)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenOwnerValidateSignatureBN(t *testing.T) {
	msg := make([]byte, 32)
	ets := &ERCTokenSignature{}
	erctp := &ERCTPreImage{}
	err := erctp.Owner.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	eto := &ERCTokenOwner{}
	err = eto.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	acct := make([]byte, constants.OwnerLen)
	eto.New(acct, constants.CurveBN256Eth)
	err = eto.ValidateSignature(msg, ets)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	bnSigner := &crypto.BNSigner{}
	privk := make([]byte, 32)
	privk[0] = 1
	privk[31] = 1
	err = bnSigner.SetPrivk(privk)
	if err != nil {
		t.Fatal(err)
	}
	secpSigner := &crypto.Secp256k1Signer{}
	if err := secpSigner.SetPrivk(privk); err != nil {
		t.Fatal(err)
	}
	etsSecp, err := eto.Sign(msg, secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.ValidateSignature(msg, etsSecp)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	msgBad := make([]byte, len(msg)+1)
	etsBad, err := eto.Sign(msgBad, bnSigner)
	if err != nil {
		t.Fatal(err)
	}
	// Sig validation will fail
	err = eto.ValidateSignature(msg, etsBad)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
	// accounts will not match
	err = eto.ValidateSignature(msgBad, etsBad)
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	// Correct vso
	pk, err := bnSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	eto.Account = crypto.GetAccount(pk)
	ets, err = eto.Sign(msg, bnSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = eto.ValidateSignature(msg, ets)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenSignatureMarshalBinary(t *testing.T) {
	ets := &ERCTokenSignature{}
	_, err := ets.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	ets.SVA = ERCTokenSVA
	_, err = ets.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	ets.CurveSpec = constants.CurveSecp256k1
	_, err = ets.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	ets.Signature = make([]byte, constants.CurveSecp256k1SigLen)
	_, err = ets.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenSignatureUnmarshalBinary(t *testing.T) {
	ets := &ERCTokenSignature{}
	signature := make([]byte, 0)
	err := ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	signature = make([]byte, 1)
	signature[0] = uint8(ERCTokenSVA)
	err = ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	signature = make([]byte, 2)
	signature[0] = uint8(ERCTokenSVA)
	signature[1] = uint8(constants.CurveSecp256k1)
	err = ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	signature = make([]byte, 2)
	signature[0] = uint8(ERCTokenSVA)
	signature[1] = uint8(constants.CurveBN256Eth)
	err = ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	signature = make([]byte, 2)
	signature[0] = uint8(ERCTokenSVA)
	err = ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	signature = make([]byte, 2+constants.CurveSecp256k1SigLen+1)
	signature[0] = uint8(ERCTokenSVA)
	signature[1] = uint8(constants.CurveSecp256k1)
	err = ets.UnmarshalBinary(signature)
	if err == nil {
		t.Fatal("Should have raised error (6)")
	}

	signature = make([]byte, 2+constants.CurveSecp256k1SigLen)
	signature[0] = uint8(ERCTokenSVA)
	signature[1] = uint8(constants.CurveSecp256k1)
	err = ets.UnmarshalBinary(signature)
	if err != nil {
		t.Fatal(err)
	}
}
