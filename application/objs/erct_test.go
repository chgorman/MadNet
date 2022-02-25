package objs

import (
	"bytes"
	"testing"

	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
)

func TestERCTokenGood(t *testing.T) {
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
	tokenID, err := new(uint256.Uint256).FromUint64(65541)
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
		Value:                val,
		Withdraw:             withdraw,
		TXOutIdx:             txoid,
		Owner:                owner,
		SmartContractAddress: sca,
		Fee:                  fee,
		TokenID:              tokenID,
	}
	txHash := make([]byte, constants.HashLen)
	erct := &ERCToken{
		ERCTPreImage: erctpi,
		TxHash:       txHash,
	}
	erct2 := &ERCToken{}
	erctBytes, err := erct.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	err = erct2.UnmarshalBinary(erctBytes)
	if err != nil {
		t.Fatal(err)
	}
	erctEqual(t, erct, erct2)
}

func erctEqual(t *testing.T, erct1, erct2 *ERCToken) {
	erctpi1 := erct1.ERCTPreImage
	erctpi2 := erct2.ERCTPreImage
	erctpiEqual(t, erctpi1, erctpi2)
	if !bytes.Equal(erct1.TxHash, erct2.TxHash) {
		t.Fatal("Do not agree on TxHash!")
	}
}

func TestERCTokenMarshalBinaryBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	_, err = utxo.ercToken.MarshalCapn(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct := &ERCToken{}
	_, err = erct.MarshalBinary()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	_, err = erct.MarshalCapn(nil)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenUnmarshalBinaryBad(t *testing.T) {
	utxo := TXOut{}
	err := utxo.ercToken.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	err = utxo.ercToken.UnmarshalCapn(mdefs.ERCToken{})
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct := &ERCToken{}
	err = erct.UnmarshalBinary(nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	err = erct.UnmarshalCapn(mdefs.ERCToken{})
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenNew(t *testing.T) {
	utxo := &TXOut{}
	err := utxo.ercToken.New(0, 0, nil, nil, nil, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	erct := &ERCToken{}
	err = erct.New(0, 0, nil, nil, nil, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	value := uint256.Zero()
	err = erct.New(0, 0, value, nil, nil, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	value = uint256.One()
	err = erct.New(0, 0, value, nil, nil, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	fee := uint256.Zero()
	err = erct.New(0, 0, value, fee, nil, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	tokenID := uint256.Zero()
	err = erct.New(0, 0, value, fee, tokenID, nil, 0, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("secret"))); err != nil {
		t.Fatal(err)
	}
	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &ValueStoreOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)
	if err != nil {
		t.Fatal(err)
	}
	err = erct.New(0, 0, value, fee, tokenID, owner.Account, owner.CurveSpec, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (6)")
	}

	chainID := uint32(1)
	err = erct.New(chainID, 0, value, fee, tokenID, owner.Account, owner.CurveSpec, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (7)")
	}

	exitChainID := uint32(2)
	err = erct.New(chainID, exitChainID, value, fee, tokenID, owner.Account, owner.CurveSpec, false, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (8)")
	}

	txhash := make([]byte, constants.HashLen)
	err = erct.New(chainID, exitChainID, value, fee, tokenID, owner.Account, owner.CurveSpec, false, txhash, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenNewFromDeposit(t *testing.T) {
	utxo := &TXOut{}
	err := utxo.ercToken.NewFromDeposit(0, 0, nil, nil, nil, 0, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (0)")
	}

	erct := &ERCToken{}
	err = erct.NewFromDeposit(0, 0, nil, nil, nil, 0, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("secret"))); err != nil {
		t.Fatal(err)
	}
	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &ValueStoreOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)
	if err != nil {
		t.Fatal(err)
	}
	err = erct.NewFromDeposit(0, 0, nil, nil, owner.Account, owner.CurveSpec, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}

	chainID := uint32(1)
	err = erct.NewFromDeposit(chainID, 0, nil, nil, owner.Account, owner.CurveSpec, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}

	exitChainID := uint32(2)
	err = erct.NewFromDeposit(chainID, exitChainID, nil, nil, owner.Account, owner.CurveSpec, nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}

	nonce := make([]byte, constants.HashLen)
	nonce[0] = 1
	err = erct.NewFromDeposit(chainID, exitChainID, nil, nil, owner.Account, owner.CurveSpec, nonce, nil)
	if err == nil {
		t.Fatal("Should have raised error (5)")
	}

	value := uint256.Zero()
	err = erct.NewFromDeposit(chainID, exitChainID, value, nil, owner.Account, owner.CurveSpec, nonce, nil)
	if err == nil {
		t.Fatal("Should have raised error (6)")
	}

	value = uint256.One()
	err = erct.NewFromDeposit(chainID, exitChainID, value, nil, owner.Account, owner.CurveSpec, nonce, nil)
	if err == nil {
		t.Fatal("Should have raised error (7)")
	}

	tokenID := uint256.Zero()
	err = erct.NewFromDeposit(chainID, exitChainID, value, tokenID, owner.Account, owner.CurveSpec, nonce, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenPreHashBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.PreHash()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.PreHash()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenUTXOIDBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.UTXOID()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.UTXOID()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenTxOutIdxBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.TxOutIdx()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.TxOutIdx()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenSetTxOutIdxBad(t *testing.T) {
	utxo := TXOut{}
	err := utxo.ercToken.SetTxOutIdx(0)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.SetTxOutIdx(0)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenSetTxHashBad(t *testing.T) {
	utxo := TXOut{}
	err := utxo.ercToken.SetTxHash(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.SetTxHash(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	err = erct.SetTxHash(nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenChainIDBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.ChainID()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.ChainID()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.ChainID()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenValueBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.Value()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.Value()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.Value()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	erct.ERCTPreImage.Value = &uint256.Uint256{}
	_, err = erct.Value()
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenFeeBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.Fee()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.Fee()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.Fee()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenTokenIDBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.TokenID()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.TokenID()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.TokenID()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenOwnerBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.Owner()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.Owner()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.Owner()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenGenericOwnerBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.GenericOwner()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.GenericOwner()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenSignBad(t *testing.T) {
	utxo := TXOut{}
	txin := &TXIn{}
	err := utxo.ercToken.Sign(txin, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.Sign(nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestERCTokenValidateSignatureBad(t *testing.T) {
	utxo := TXOut{}
	err := utxo.ercToken.ValidateSignature(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.ValidateSignature(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	txIn := &TXIn{}
	err = erct.ValidateSignature(txIn)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenMakeTxInBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.MakeTxIn()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.MakeTxIn()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.MakeTxIn()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	erct.ERCTPreImage.ChainID = 1
	_, err = erct.MakeTxIn()
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}
