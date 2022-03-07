package objs

import (
	"bytes"
	"math/big"
	"testing"

	mdefs "github.com/MadBase/MadNet/application/objs/capn"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/MadBase/MadNet/utils"
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

func TestERCTokenPreHashGood(t *testing.T) {
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
	withdraw := false

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
	erct := &ERCToken{
		ERCTPreImage: erctpi,
	}

	preHash, err := erct.PreHash()
	if err != nil {
		t.Fatal(err)
	}

	preHash2, err := erct.PreHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(preHash, preHash2) {
		t.Fatal("PreHash values do not match (1)")
	}
	if !bytes.Equal(preHash, erct.ERCTPreImage.preHash) {
		t.Fatal("PreHash values do not match (2)")
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
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.UTXOID()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
}

func TestERCTokenUTXOIDGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	txhash := crypto.Hasher([]byte("txhash"))
	trueUTXOID := make([]byte, constants.HashLen)
	trueUTXOID[0] = 1
	trueUTXOID[1] = 2
	erct.TxHash = txhash
	erct.utxoID = utils.CopySlice(trueUTXOID)
	utxoID, err := erct.UTXOID()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(trueUTXOID, utxoID) {
		t.Fatal("utxoIDs do not math (1)")
	}
	utxoID2, err := erct.UTXOID()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(trueUTXOID, utxoID2) {
		t.Fatal("utxoIDs do not math (2)")
	}
}

func TestERCTokenUTXOIDGood2(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	txhash := crypto.Hasher([]byte("txhash"))
	erct.TxHash = txhash
	trueUtxoID, err := erct.UTXOID()
	if err != nil {
		t.Fatal(err)
	}
	utxoID, err := erct.UTXOID()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(trueUtxoID, utxoID) {
		t.Fatal("utxoIDs do not math")
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

func TestERCTokenSetTxOutIdxGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	idx := uint32(17)
	err := erct.SetTxOutIdx(idx)
	if err != nil {
		t.Fatal(err)
	}
	txOutIdx, err := erct.TxOutIdx()
	if err != nil {
		t.Fatal(err)
	}
	if txOutIdx != idx {
		t.Fatal("invalid TxOutIdx")
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

func TestERCTokenSetTxHashGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	txhash := crypto.Hasher([]byte("txhash"))
	err := erct.SetTxHash(txhash)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(txhash, erct.TxHash) {
		t.Fatal("invalid txhash")
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

func TestERCTokenERCValueBad(t *testing.T) {
	utxo := TXOut{}
	_, err := utxo.ercToken.ERCValue()
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	_, err = erct.ERCValue()
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	_, err = erct.ERCValue()
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	erct.ERCTPreImage.Value = &uint256.Uint256{}
	_, err = erct.ERCValue()
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenERCValueGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	val := uint256.Two()
	erct.ERCTPreImage.Value = val.Clone()
	retValue, err := erct.ERCValue()
	if err != nil {
		t.Fatal(err)
	}
	if !retValue.Eq(val) {
		t.Fatal("incorrect ERC Value")
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

func TestERCTokenFeeGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	fee := uint256.Two()
	erct.ERCTPreImage.Fee = fee.Clone()
	retFee, err := erct.Fee()
	if err != nil {
		t.Fatal(err)
	}
	if !retFee.Eq(fee) {
		t.Fatal("invalid fee")
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

func TestERCTokenTokenIDGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	tokenID := uint256.Two()
	erct.ERCTPreImage.TokenID = tokenID.Clone()
	retTokenID, err := erct.TokenID()
	if err != nil {
		t.Fatal(err)
	}
	if !tokenID.Eq(retTokenID) {
		t.Fatal("invalid tokenID")
	}
}

func TestERCTokenIsDeposit(t *testing.T) {
	utxo := TXOut{}
	isDeposit := utxo.ercToken.IsDeposit()
	if isDeposit {
		t.Fatal("should be false (1)")
	}
	erct := &ERCToken{}
	isDeposit = erct.IsDeposit()
	if isDeposit {
		t.Fatal("should be false (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	isDeposit = erct.IsDeposit()
	if isDeposit {
		t.Fatal("should be false (3)")
	}
	erct.ERCTPreImage.TXOutIdx = constants.MaxUint32
	isDeposit = erct.IsDeposit()
	if !isDeposit {
		t.Fatal("should be true")
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

func TestERCTokenOwnerGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	acct := crypto.Hasher([]byte("ercowner"))[12:]
	ercowner := &ERCTokenOwner{}
	curveSpec := constants.CurveSecp256k1
	ercowner.New(acct, curveSpec)
	ercoBytes, err := ercowner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	erco2 := &ERCTokenOwner{}
	err = erco2.UnmarshalBinary(ercoBytes)
	if err != nil {
		t.Fatal(err)
	}
	erct.ERCTPreImage.Owner = erco2
	retOwner, err := erct.Owner()
	if err != nil {
		t.Fatal(err)
	}
	retOwnerBytes, err := retOwner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(retOwnerBytes, ercoBytes) {
		t.Fatal("invalid erc owner")
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

func TestERCTokenGenericOwnerGood(t *testing.T) {
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	acct := crypto.Hasher([]byte("ercowner"))[12:]
	ercowner := &ERCTokenOwner{}
	curveSpec := constants.CurveSecp256k1
	ercowner.New(acct, curveSpec)
	owner := &Owner{}
	err := owner.NewFromERCTokenOwner(ercowner)
	if err != nil {
		t.Fatal(err)
	}
	ownerBytes, err := owner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	etoBytes, err := ercowner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	eto2 := &ERCTokenOwner{}
	err = eto2.UnmarshalBinary(etoBytes)
	if err != nil {
		t.Fatal(err)
	}
	erct.ERCTPreImage.Owner = eto2
	retOwner, err := erct.GenericOwner()
	if err != nil {
		t.Fatal(err)
	}
	retOwnerBytes, err := retOwner.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ownerBytes, retOwnerBytes) {
		t.Fatal("invalid generic owner")
	}
}

func TestERCTokenSignBad(t *testing.T) {
	utxo := TXOut{}
	txIn := &TXIn{}
	err := utxo.ercToken.Sign(txIn, nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct := &ERCToken{}
	err = erct.Sign(nil, nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	txIn.TXInLinker = &TXInLinker{}
	txIn.TXInLinker.TXInPreImage = &TXInPreImage{}
	cid := uint32(1)
	consumedTxOutIdx := uint32(13)
	consumedTxHash := crypto.Hasher([]byte("ConsumedTxHash"))
	txIn.TXInLinker.TXInPreImage.ChainID = cid
	txIn.TXInLinker.TXInPreImage.ConsumedTxIdx = consumedTxOutIdx
	txIn.TXInLinker.TXInPreImage.ConsumedTxHash = consumedTxHash
	err = erct.Sign(txIn, nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	acct := crypto.Hasher([]byte("ercowner"))[12:]
	ercowner := &ERCTokenOwner{}
	curveSpec := constants.CurveSecp256k1
	ercowner.New(acct, curveSpec)
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.Owner = ercowner
	err = erct.Sign(txIn, nil)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenSignGood(t *testing.T) {
	txIn := &TXIn{}
	txIn.TXInLinker = &TXInLinker{}
	txIn.TXInLinker.TXInPreImage = &TXInPreImage{}
	cid := uint32(1)
	consumedTxOutIdx := uint32(13)
	consumedTxHash := crypto.Hasher([]byte("ConsumedTxHash"))
	txIn.TXInLinker.TXInPreImage.ChainID = cid
	txIn.TXInLinker.TXInPreImage.ConsumedTxIdx = consumedTxOutIdx
	txIn.TXInLinker.TXInPreImage.ConsumedTxHash = consumedTxHash
	erct := &ERCToken{}
	secpSigner := &crypto.Secp256k1Signer{}
	privk := make([]byte, 32)
	privk[0] = 1
	privk[31] = 1
	if err := secpSigner.SetPrivk(privk); err != nil {
		t.Fatal(err)
	}
	pubkey, err := secpSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	acct := crypto.GetAccount(pubkey)
	ercowner := &ERCTokenOwner{}
	curveSpec := constants.CurveSecp256k1
	ercowner.New(acct, curveSpec)
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.Owner = ercowner
	err = erct.Sign(txIn, secpSigner)
	if err != nil {
		t.Fatal(err)
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
	txIn.TXInLinker = &TXInLinker{}
	txIn.TXInLinker.TXInPreImage = &TXInPreImage{}
	cid := uint32(1)
	consumedTxOutIdx := uint32(13)
	consumedTxHash := crypto.Hasher([]byte("ConsumedTxHash"))
	txIn.TXInLinker.TXInPreImage.ChainID = cid
	txIn.TXInLinker.TXInPreImage.ConsumedTxIdx = consumedTxOutIdx
	txIn.TXInLinker.TXInPreImage.ConsumedTxHash = consumedTxHash
	err = erct.ValidateSignature(txIn)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenValidateSignatureGood(t *testing.T) {
	txIn := &TXIn{}
	txIn.TXInLinker = &TXInLinker{}
	txIn.TXInLinker.TXInPreImage = &TXInPreImage{}
	cid := uint32(1)
	consumedTxOutIdx := uint32(13)
	consumedTxHash := crypto.Hasher([]byte("ConsumedTxHash"))
	txIn.TXInLinker.TXInPreImage.ChainID = cid
	txIn.TXInLinker.TXInPreImage.ConsumedTxIdx = consumedTxOutIdx
	txIn.TXInLinker.TXInPreImage.ConsumedTxHash = consumedTxHash
	erct := &ERCToken{}
	secpSigner := &crypto.Secp256k1Signer{}
	privk := make([]byte, 32)
	privk[0] = 1
	privk[31] = 1
	if err := secpSigner.SetPrivk(privk); err != nil {
		t.Fatal(err)
	}
	pubkey, err := secpSigner.Pubkey()
	if err != nil {
		t.Fatal(err)
	}
	acct := crypto.GetAccount(pubkey)
	ercowner := &ERCTokenOwner{}
	curveSpec := constants.CurveSecp256k1
	ercowner.New(acct, curveSpec)
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.Owner = ercowner
	err = erct.Sign(txIn, secpSigner)
	if err != nil {
		t.Fatal(err)
	}
	err = erct.ValidateSignature(txIn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestERCTokenValidateFeeBad(t *testing.T) {
	erct := &ERCToken{}
	err := erct.ValidateFee(nil)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.Fee = uint256.One()
	erct.ERCTPreImage.TXOutIdx = constants.MaxUint32
	err = erct.ValidateFee(nil)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
	erct.ERCTPreImage.TXOutIdx = 0
	err = erct.ValidateFee(nil)
	if err == nil {
		t.Fatal("Should have raised error (3)")
	}
	msg := makeMockStorageGetter()
	storage := makeStorage(msg)
	err = erct.ValidateFee(storage)
	if err == nil {
		t.Fatal("Should have raised error (4)")
	}
}

func TestERCTokenValidateFeeGood(t *testing.T) {
	msg := makeMockStorageGetter()
	ercFee := uint256.One()
	msg.SetERCTokenFee(big.NewInt(1))
	storage := makeStorage(msg)
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.Fee = uint256.Zero()
	erct.ERCTPreImage.TXOutIdx = constants.MaxUint32
	err := erct.ValidateFee(storage)
	if err != nil {
		t.Fatal(err)
	}
	erct.ERCTPreImage.Fee = ercFee.Clone()
	erct.ERCTPreImage.TXOutIdx = 0
	err = erct.ValidateFee(storage)
	if err != nil {
		t.Fatal(err)
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

func TestERCTokenMakeTxInGood(t *testing.T) {
	cid := uint32(1)
	txOutIdx := uint32(2)
	txhash := crypto.Hasher([]byte("TxHash"))
	erct := &ERCToken{}
	erct.ERCTPreImage = &ERCTPreImage{}
	erct.ERCTPreImage.ChainID = cid
	erct.ERCTPreImage.TXOutIdx = txOutIdx
	erct.TxHash = txhash
	txIn, err := erct.MakeTxIn()
	if err != nil {
		t.Fatal(err)
	}
	retCid, err := txIn.ChainID()
	if err != nil {
		t.Fatal(err)
	}
	if retCid != cid {
		t.Fatal("invalid chainID")
	}
	retTxOutIdx, err := txIn.ConsumedTxIdx()
	if err != nil {
		t.Fatal(err)
	}
	if retTxOutIdx != txOutIdx {
		t.Fatal("invalid txOutIdx")
	}
	retConsumedTxHash, err := txIn.ConsumedTxHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(retConsumedTxHash, txhash) {
		t.Fatal("invalid consumedTxHash")
	}
}

func TestERCTokenWithdrawTokens(t *testing.T) {
	utxo := TXOut{}
	retWithdraw := utxo.ercToken.WithdrawTokens()
	if retWithdraw {
		t.Fatal("Should not withdraw tokens (1)")
	}
	erct := &ERCToken{}
	retWithdraw = erct.WithdrawTokens()
	if retWithdraw {
		t.Fatal("Should not withdraw tokens (2)")
	}
	erct.ERCTPreImage = &ERCTPreImage{}
	retWithdraw = erct.WithdrawTokens()
	if retWithdraw {
		t.Fatal("Should not withdraw tokens (3)")
	}
	erct.ERCTPreImage.Withdraw = true
	retWithdraw = erct.WithdrawTokens()
	if !retWithdraw {
		t.Fatal("Should withdraw tokens")
	}
}
