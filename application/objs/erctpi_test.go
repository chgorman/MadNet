package objs

import (
	"bytes"
	"testing"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/utils"
)

func TestERCTPreImageGood(t *testing.T) {
	cid := uint32(2)
	ecid := uint32(3)
	val, err := new(uint256.Uint256).FromUint64(65537)
	if err != nil {
		t.Fatal(err)
	}
	fee, err := new(uint256.Uint256).FromUint64(65539)
	if err != nil {
		t.Fatal(err)
	}
	tid, err := new(uint256.Uint256).FromUint64(65541)
	if err != nil {
		t.Fatal(err)
	}
	txoid := uint32(17)
	withdraw := true

	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("a"))); err != nil {
		t.Fatal(err)
	}
	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &ERCTokenOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)

	scaBytes := crypto.GetAccount(crypto.Hasher([]byte("SmartContractAddress")))
	sca := &SmartContract{}
	err = sca.UnmarshalBinary(scaBytes)
	if err != nil {
		t.Fatal(err)
	}

	erctpi := &ERCTPreImage{
		ChainID:              cid,
		ExitChainID:          ecid,
		TXOutIdx:             txoid,
		Withdraw:             withdraw,
		SmartContractAddress: sca,
		Owner:                owner,
		Value:                val,
		TokenID:              tid,
		Fee:                  fee,
	}
	erctpi2 := &ERCTPreImage{}
	erctpiBytes, err := erctpi.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	err = erctpi2.UnmarshalBinary(erctpiBytes)
	if err != nil {
		t.Fatal(err)
	}
	erctpiEqual(t, erctpi, erctpi2)
}

func erctpiEqual(t *testing.T, erctpi1, erctpi2 *ERCTPreImage) {
	if erctpi1.ChainID != erctpi2.ChainID {
		t.Fatal("Do not agree on ChainID!")
	}
	if erctpi1.ExitChainID != erctpi2.ExitChainID {
		t.Fatal("Do not agree on ExitChainID!")
	}
	if erctpi1.Withdraw != erctpi2.Withdraw {
		t.Fatal("Do not agree on Withdraw!")
	}
	if !erctpi1.Value.Eq(erctpi2.Value) {
		t.Fatal("Do not agree on Value!")
	}
	if !erctpi1.Fee.Eq(erctpi2.Fee) {
		t.Fatal("Do not agree on Fee!")
	}
	if !erctpi1.TokenID.Eq(erctpi2.TokenID) {
		t.Fatal("Do not agree on TokenID!")
	}
	if erctpi1.TXOutIdx != erctpi2.TXOutIdx {
		t.Fatal("Do not agree on TXOutIdx!")
	}
	if !bytes.Equal(erctpi1.Owner.Account, erctpi2.Owner.Account) {
		t.Fatal("Do not agree on Owner!")
	}
	if !bytes.Equal(erctpi1.SmartContractAddress.address, erctpi2.SmartContractAddress.address) {
		t.Fatal("Do not agree on SmartContractAddress!")
	}
}

func TestERCTPreImageUnmarshalBad1(t *testing.T) {
	erctpi := &ERCTPreImage{}
	err := erctpi.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.ERCTPreImage.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTPreImageMarshalBad1(t *testing.T) {
	erctpi := &ERCTPreImage{}
	_, err := erctpi.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.ERCTPreImage.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	_, err = erct.ERCTPreImage.MarshalCapn(nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTPreImageMarshalBad2(t *testing.T) {
	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("a"))); err != nil {
		t.Fatal(err)
	}
	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &ERCTokenOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)

	erctpi := &ERCTPreImage{
		Owner: owner,
	}
	_, err = erctpi.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error")
	}
}

func TestERCTPreImagePreHashBad(t *testing.T) {
	erct := &ERCToken{}
	_, err := erct.ERCTPreImage.PreHash()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erctpi := &ERCTPreImage{}
	_, err = erctpi.PreHash()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTPreImagePreHashGood(t *testing.T) {
	cid := uint32(2)
	ecid := uint32(3)
	val, err := new(uint256.Uint256).FromUint64(65537)
	if err != nil {
		t.Fatal(err)
	}
	fee, err := new(uint256.Uint256).FromUint64(65539)
	if err != nil {
		t.Fatal(err)
	}
	tid, err := new(uint256.Uint256).FromUint64(65541)
	if err != nil {
		t.Fatal(err)
	}
	txoid := uint32(17)
	withdraw := true

	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("a"))); err != nil {
		t.Fatal(err)
	}
	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &ERCTokenOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)

	scaBytes := crypto.GetAccount(crypto.Hasher([]byte("SmartContractAddress")))
	sca := &SmartContract{}
	err = sca.UnmarshalBinary(scaBytes)
	if err != nil {
		t.Fatal(err)
	}

	erctpi := &ERCTPreImage{
		ChainID:              cid,
		ExitChainID:          ecid,
		TXOutIdx:             txoid,
		Withdraw:             withdraw,
		SmartContractAddress: sca,
		Owner:                owner,
		Value:                val,
		TokenID:              tid,
		Fee:                  fee,
	}
	ret, err := erctpi.PreHash()
	if err != nil {
		t.Fatal(err)
	}
	ret2, err := erctpi.PreHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ret, ret2) {
		t.Fatal("PreHash should be the same")
	}
}

func TestERCTPreImageValidateSignatureBad(t *testing.T) {
	erct := &ERCToken{}
	err := erct.ERCTPreImage.ValidateSignature(nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erctpi := &ERCTPreImage{}
	err = erctpi.ValidateSignature(nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTPreImageValidateTokenBad(t *testing.T) {
	erct := &ERCToken{}
	err := erct.ERCTPreImage.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	erctpi := &ERCTPreImage{}
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	cid := uint32(1)
	erctpi.ChainID = cid
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	ecid := uint32(2)
	erctpi.ExitChainID = ecid
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	eto := &ERCTokenOwner{}
	acct := make([]byte, constants.OwnerLen)
	eto.New(acct, constants.CurveSecp256k1)
	erctpi.Owner = eto
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	value := uint256.Zero()
	erctpi.Value = value
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (6)")
	}

	value = uint256.One()
	erctpi.Value = value
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (7)")
	}

	fee := uint256.Zero()
	erctpi.Fee = fee
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (8)")
	}

	tokenID := uint256.Two()
	erctpi.TokenID = tokenID
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (9)")
	}

	sca := &SmartContract{}
	erctpi.SmartContractAddress = sca
	err = erctpi.ValidateToken()
	if err == nil {
		t.Fatal("Should have raised error (10)")
	}
}

func TestERCTPreImageValidateTokenGood(t *testing.T) {
	erctpi := &ERCTPreImage{}
	cid := uint32(1)
	erctpi.ChainID = cid
	ecid := uint32(2)
	erctpi.ExitChainID = ecid
	eto := &ERCTokenOwner{}
	acct := make([]byte, constants.OwnerLen)
	eto.New(acct, constants.CurveSecp256k1)
	erctpi.Owner = eto
	value := uint256.One()
	erctpi.Value = value
	fee := uint256.Zero()
	erctpi.Fee = fee
	tokenID := uint256.Two()
	erctpi.TokenID = tokenID
	addr := make([]byte, constants.OwnerLen)
	addr[0] = 1
	sca := &SmartContract{}
	err := sca.UnmarshalBinary(addr)
	if err != nil {
		t.Fatal(err)
	}
	erctpi.SmartContractAddress = sca
	err = erctpi.ValidateToken()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSmartContractNew(t *testing.T) {
	erctpi := &ERCTPreImage{}
	err := erctpi.SmartContractAddress.New(nil)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}
	sc := &SmartContract{}
	err = sc.New(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	data := make([]byte, constants.OwnerLen)
	data[0] = 1
	err = sc.New(data)
	if err != nil {
		t.Fatal(err)
	}
	retData, err := sc.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, retData) {
		t.Fatal("invalid sc")
	}
}

func TestSmartContractMarshalBad(t *testing.T) {
	erctpi := &ERCTPreImage{}
	_, err := erctpi.SmartContractAddress.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	sc := &SmartContract{}
	_, err = sc.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestSmartContractMarshalGood(t *testing.T) {
	sca := make([]byte, constants.OwnerLen)
	for k := 0; k < len(sca); k++ {
		sca[k] = byte(k + 1)
	}
	sc := &SmartContract{}
	sc.address = utils.CopySlice(sca)
	ret, err := sc.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(sca, ret) {
		t.Fatal("invalid address")
	}
}

func TestSmartContractUnmarshalBad(t *testing.T) {
	erctpi := &ERCTPreImage{}
	err := erctpi.SmartContractAddress.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	sc := &SmartContract{}
	err = sc.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestSmartContractUnmarshalGood(t *testing.T) {
	sca := make([]byte, constants.OwnerLen)
	for k := 0; k < len(sca); k++ {
		sca[k] = byte(k + 1)
	}
	sc := &SmartContract{}
	err := sc.UnmarshalBinary(sca)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(sc.address, sca) {
		t.Fatal("invalid address")
	}
}
