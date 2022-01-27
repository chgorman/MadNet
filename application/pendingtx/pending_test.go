package pendingtx

import (
	"bytes"
	"context"
	"crypto/rand"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/MadBase/MadNet/application/objs"
	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/constants"
	"github.com/MadBase/MadNet/crypto"
	"github.com/dgraph-io/badger/v2"
)

type mockTrie struct {
	m map[string]bool
}

func (mt *mockTrie) IsValid(txn *badger.Txn, txs objs.TxVec, currentHeight uint32, deposits objs.Vout) (objs.Vout, error) {
	utxo1 := makeDS()
	vout := objs.Vout{utxo1}
	return vout, nil
}

func (mt *mockTrie) TrieContains(txn *badger.Txn, utxo []byte) (bool, error) {
	return mt.m[string(utxo)], nil
}

func (mt *mockTrie) Add(utxo []byte) {
	mt.m[string(utxo)] = true
}

func (mt *mockTrie) Remove(utxo []byte) {
	delete(mt.m, string(utxo))
}

func (mt *mockTrie) Get(txn *badger.Txn, utxoIDs [][]byte) ([]*objs.TXOut, [][]byte, []*objs.TXOut, error) {
	return nil, nil, nil, nil
}

func testingOwner() objs.Signer {
	signer := &crypto.Secp256k1Signer{}
	err := signer.SetPrivk(crypto.Hasher([]byte("secret")))
	if err != nil {
		panic(err)
	}
	return signer
}

func accountFromSigner(s objs.Signer) []byte {
	pubk, err := s.Pubkey()
	if err != nil {
		panic(err)
	}
	return crypto.GetAccount(pubk)
}

func makeVS(ownerSigner objs.Signer, fee *uint256.Uint256) *objs.TXOut {
	cid := uint32(2)
	val := uint256.One()

	ownerAcct := accountFromSigner(ownerSigner)
	owner := &objs.ValueStoreOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)

	vsp := &objs.VSPreImage{
		ChainID: cid,
		Value:   val,
		Owner:   owner,
		Fee:     fee,
	}
	vs := &objs.ValueStore{
		VSPreImage: vsp,
		TxHash:     make([]byte, constants.HashLen),
	}
	utxInputs := &objs.TXOut{}
	err := utxInputs.NewValueStore(vs)
	if err != nil {
		panic(err)
	}
	return utxInputs
}

func makeVSTXIn(ownerSigner objs.Signer, txHash []byte) (*objs.TXOut, *objs.TXIn) {
	vs := makeVS(ownerSigner, uint256.One())
	vss, err := vs.ValueStore()
	if err != nil {
		panic(err)
	}
	if txHash == nil {
		txHash = make([]byte, constants.HashLen)
		_, err := rand.Read(txHash)
		if err != nil {
			panic(err)
		}
	}
	vss.TxHash = txHash
	txin, err := vss.MakeTxIn()
	if err != nil {
		panic(err)
	}
	return vs, txin
}

func makeTxInitial() (objs.Vout, *objs.Tx) {
	ownerSigner := testingOwner()
	consumedUTXOs := objs.Vout{}
	txInputs := []*objs.TXIn{}
	for i := 0; i < 2; i++ {
		utxo, txin := makeVSTXIn(ownerSigner, nil)
		consumedUTXOs = append(consumedUTXOs, utxo)
		txInputs = append(txInputs, txin)
	}
	generatedUTXOs := objs.Vout{}
	for i := 0; i < 2; i++ {
		generatedUTXOs = append(generatedUTXOs, makeVS(ownerSigner, uint256.One()))
	}
	err := generatedUTXOs.SetTxOutIdx()
	if err != nil {
		panic(err)
	}
	txfee := uint256.Zero()
	tx := &objs.Tx{
		Vin:  txInputs,
		Vout: generatedUTXOs,
		Fee:  txfee,
	}
	err = tx.SetTxHash()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		vs, err := consumedUTXOs[i].ValueStore()
		if err != nil {
			panic(err)
		}
		err = vs.Sign(tx.Vin[i], ownerSigner)
		if err != nil {
			panic(err)
		}
	}
	return consumedUTXOs, tx
}

func makeTxConsuming(consumedUTXOs objs.Vout) *objs.Tx {
	ownerSigner := testingOwner()
	txInputs := []*objs.TXIn{}
	for i := 0; i < 2; i++ {
		txin, err := consumedUTXOs[i].MakeTxIn()
		if err != nil {
			panic(err)
		}
		txInputs = append(txInputs, txin)
	}
	generatedUTXOs := objs.Vout{}
	for i := 0; i < 2; i++ {
		generatedUTXOs = append(generatedUTXOs, makeVS(ownerSigner, uint256.One()))
	}
	err := generatedUTXOs.SetTxOutIdx()
	if err != nil {
		panic(err)
	}
	txfee := uint256.Zero()
	tx := &objs.Tx{
		Vin:  txInputs,
		Vout: generatedUTXOs,
		Fee:  txfee,
	}
	err = tx.SetTxHash()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		vs, err := consumedUTXOs[i].ValueStore()
		if err != nil {
			panic(err)
		}
		err = vs.Sign(tx.Vin[i], ownerSigner)
		if err != nil {
			panic(err)
		}
	}
	return tx
}

func makeTxCleanup() (objs.Vout, *objs.Tx) {
	ownerSigner := testingOwner()
	consumedUTXOs := objs.Vout{}
	txInputs := []*objs.TXIn{}
	utxo, txin := makeVSTXIn(ownerSigner, nil)
	consumedUTXOs = append(consumedUTXOs, utxo)
	txInputs = append(txInputs, txin)

	generatedUTXOs := objs.Vout{}
	generatedUTXOs = append(generatedUTXOs, makeVS(ownerSigner, uint256.Zero()))
	err := generatedUTXOs.SetTxOutIdx()
	if err != nil {
		panic(err)
	}
	txfee := uint256.Zero()
	tx := &objs.Tx{
		Vin:  txInputs,
		Vout: generatedUTXOs,
		Fee:  txfee,
	}
	err = tx.SetTxHash()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1; i++ {
		vs, err := consumedUTXOs[i].ValueStore()
		if err != nil {
			panic(err)
		}
		err = vs.Sign(tx.Vin[i], ownerSigner)
		if err != nil {
			panic(err)
		}
	}
	return consumedUTXOs, tx
}

func makeDS() *objs.TXOut {
	ownerSigner := &crypto.Secp256k1Signer{}
	if err := ownerSigner.SetPrivk(crypto.Hasher([]byte("a"))); err != nil {
		panic(err)
	}
	fee := uint256.Zero()
	rawData1 := make([]byte, constants.MaxDataStoreSize)
	iat := uint32(1)
	numEpochs := uint32(1)
	index1 := make([]byte, constants.HashLen)
	index1[0] = 1
	return makeDSWithValueFee(ownerSigner, 0, rawData1, index1, iat, numEpochs, fee)
}

func makeDSWithValueFee(ownerSigner objs.Signer, i int, rawData []byte, index []byte, startEpoch uint32, numEpochs uint32, dsPerEpochFee *uint256.Uint256) *objs.TXOut {
	if dsPerEpochFee == nil || len(rawData) == 0 {
		panic("invalid fee or rawData")
	}
	cid := uint32(2)

	ownerPubk, err := ownerSigner.Pubkey()
	if err != nil {
		panic(err)
	}
	ownerAcct := crypto.GetAccount(ownerPubk)
	owner := &objs.DataStoreOwner{}
	owner.New(ownerAcct, constants.CurveSecp256k1)

	dataSize32 := uint32(len(rawData))
	deposit, err := objs.BaseDepositEquation(dataSize32, numEpochs)
	if err != nil {
		panic(err)
	}

	mulFactor, err := new(uint256.Uint256).FromUint64(uint64(numEpochs + 2))
	if err != nil {
		panic(err)
	}
	dsfee, err := new(uint256.Uint256).Mul(dsPerEpochFee, mulFactor)
	if err != nil {
		panic(err)
	}
	dsp := &objs.DSPreImage{
		ChainID:  cid,
		Index:    index,
		IssuedAt: startEpoch,
		Deposit:  deposit,
		RawData:  rawData,
		Owner:    owner,
		Fee:      dsfee,
	}
	err = dsp.ValidateDeposit()
	if err != nil {
		panic(err)
	}
	var txHash []byte
	if i == 0 {
		txHash = make([]byte, constants.HashLen)
	} else {
		txHash = crypto.Hasher([]byte(strconv.Itoa(i)))
	}
	dsl := &objs.DSLinker{
		DSPreImage: dsp,
		TxHash:     txHash,
	}
	ds := &objs.DataStore{
		DSLinker: dsl,
	}
	err = ds.PreSign(ownerSigner)
	if err != nil {
		panic(err)
	}
	err = ds.ValidatePreSignature()
	if err != nil {
		panic(err)
	}

	ds2 := &objs.DataStore{}
	dsBytes, err := ds.MarshalBinary()
	if err != nil {
		panic(err)
	}
	err = ds2.UnmarshalBinary(dsBytes)
	if err != nil {
		panic(err)
	}
	utxInputs := &objs.TXOut{}
	err = utxInputs.NewDataStore(ds)
	if err != nil {
		panic(err)
	}
	return utxInputs
}

func mustAddTx(t *testing.T, hndlr *Handler, tx *objs.Tx, currentHeight uint32) {
	err := hndlr.Add(nil, []*objs.Tx{tx}, currentHeight)
	if err != nil {
		t.Fatal(err)
	}
	mustContain(t, hndlr, tx)
}

func mustNotAdd(t *testing.T, hndlr *Handler, tx *objs.Tx, currentHeight uint32) {
	hndlr.Add(nil, []*objs.Tx{tx}, currentHeight)
	mustNotContain(t, hndlr, tx)
}

func mustContain(t *testing.T, hndlr *Handler, tx *objs.Tx) {
	txHash, err := tx.TxHash()
	if err != nil {
		t.Fatal(err)
	}
	getTx1, missing, err := hndlr.Get(nil, 1, [][]byte{txHash})
	if err != nil {
		t.Fatal(err)
	}
	if len(missing) != 0 {
		t.Fatalf("missing %x", txHash)
	}
	getTxHash1, err := getTx1[0].TxHash()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(getTxHash1, txHash) {
		t.Fatalf("txHash mismatch:\noriginalHash:%x\nreturnedHash:%x\n", txHash, getTxHash1)
	}
}

func mustNotContain(t *testing.T, hndlr *Handler, tx *objs.Tx) {
	txHash, err := tx.TxHash()
	if err != nil {
		t.Fatal(err)
	}
	_, missing, err := hndlr.Get(nil, 1, [][]byte{txHash})
	if err != nil {
		t.Fatal(err)
	}
	if len(missing) != 1 {
		t.Fatalf("delete failure: %x", txHash)
	}
	missing, err = hndlr.Contains(nil, 1, [][]byte{txHash})
	if err != nil {
		t.Fatal(err)
	}
	if len(missing) != 1 {
		t.Fatal("contains")
	}
}

func mustDelTx(t *testing.T, hndlr *Handler, tx *objs.Tx) {
	txHash, err := tx.TxHash()
	if err != nil {
		t.Fatal(err)
	}
	err = hndlr.Delete(nil, [][]byte{txHash})
	if err != nil {
		t.Fatal(err)
	}
	_, missing, err := hndlr.Get(nil, 1, [][]byte{txHash})
	if len(missing) != 1 {
		t.Fatalf("delete failure: %v", err)
	}
}

func setup(t *testing.T) (*Handler, *mockTrie, func()) {
	dir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}
	opts := badger.DefaultOptions(dir)
	db, err := badger.Open(opts)
	if err != nil {
		t.Fatal(err)
	}
	fn := func() {
		defer os.RemoveAll(dir)
		defer db.Close()
	}
	////////////////////////////////////////
	mt := &mockTrie{}
	mt.m = make(map[string]bool)
	queueSize := 1024
	hndlr, err := NewPendingTxHandler(db, queueSize)
	if err != nil {
		t.Fatal(err)
	}
	hndlr.UTXOHandler = mt
	hndlr.DepositHandler = mt
	return hndlr, mt, fn
}

func TestAdd(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()
	_, tx := makeTxInitial()
	mustAddTx(t, hndlr, tx, 1)
}

func TestAddErrors(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()
	_, tx := makeTxInitial()
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	// Attempt to add empty tx
	txBad0 := &objs.Tx{}
	err = hndlr.Add(nil, []*objs.Tx{txBad0}, 1)
	if err == nil {
		t.Fatal("Should have raised error (1)")
	}

	// Attempt to add tx with tx.Vout == nil
	txBad1 := &objs.Tx{}
	err = txBad1.UnmarshalBinary(txBytes)
	if err != nil {
		t.Fatal(err)
	}
	txBad1.Vout = nil
	err = hndlr.Add(nil, []*objs.Tx{txBad1}, 1)
	if err == nil {
		t.Fatal("Should have raised error (2)")
	}
}

func TestDel(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()
	_, tx := makeTxInitial()
	mustAddTx(t, hndlr, tx, 1)
	mustDelTx(t, hndlr, tx)
}

func TestDeleteMined(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()
	vout, tx := makeTxInitial()
	mustAddTx(t, hndlr, tx, 1)
	tx2 := makeTxConsuming(vout)
	mustAddTx(t, hndlr, tx2, 1)
	txHash, err := tx.TxHash()
	if err != nil {
		t.Fatal(err)
	}
	consumedUTXOIDs, err := tx.ConsumedUTXOID()
	if err != nil {
		t.Fatal(err)
	}
	err = hndlr.DeleteMined(nil, 1, [][]byte{txHash}, consumedUTXOIDs)
	if err != nil {
		t.Fatal(err)
	}
	mustNotContain(t, hndlr, tx2)
	mustNotAdd(t, hndlr, tx, 1)
}

func TestMissing(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()
	_, tx := makeTxInitial()
	mustAddTx(t, hndlr, tx, 1)
	_, tx2 := makeTxInitial()
	mustNotContain(t, hndlr, tx2)
}

func TestGetProposal(t *testing.T) {
	hndlr, trie, cleanup := setup(t)
	defer cleanup()
	c1, tx1 := makeTxInitial()
	mustAddTx(t, hndlr, tx1, 1)
	c2, tx2 := makeTxInitial()
	mustAddTx(t, hndlr, tx2, 1)
	tx3 := makeTxConsuming(c1)
	mustAddTx(t, hndlr, tx3, 1)
	tx4 := makeTxConsuming(c2)
	mustAddTx(t, hndlr, tx4, 1)
	maxBytes := constants.MaxUint32
	txs, _, err := hndlr.GetTxsForProposal(hndlr.db.NewTransaction(false), context.TODO(), 1, maxBytes, nil)
	if err != nil {
		t.Fatal(err)
	}
	txHashes, err := objs.TxVec(txs).TxHash()
	if err != nil {
		t.Fatal(err)
	}
	// trie must contain utxos getting spent but must not contain
	// utxos being generated
	utxoIDs, err := objs.TxVec{tx1, tx2, tx3, tx4}.ConsumedUTXOID()
	if err != nil {
		t.Fatal(err)
	}
	for _, ut := range utxoIDs {
		trie.Add(ut)
	}
	utxoIDs, err = objs.TxVec{tx1, tx2, tx3, tx4}.GeneratedUTXOID()
	if err != nil {
		t.Fatal(err)
	}
	for _, ut := range utxoIDs {
		trie.Remove(ut)
	}
	txs, err = hndlr.GetTxsForGossip(nil, context.Background(), 1, constants.MaxUint32)
	if err != nil {
		t.Fatal(err)
	}
	if len(txs) != 2 {
		t.Fatalf("conflict: %x", txHashes)
	}
}

func TestAddTxsToQueueFullQueue(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()

	_, tx := makeTxInitial()
	mustAddTx(t, hndlr, tx, 1)

	_, tx2 := makeTxInitial()
	mustAddTx(t, hndlr, tx2, 1)

	hndlr.SetQueueSize(1)
	err := hndlr.AddTxsToQueue(nil, context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddTxsToQueueCleanupTx(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()

	_, tx := makeTxCleanup()
	mustAddTx(t, hndlr, tx, 1)

	err := hndlr.AddTxsToQueue(nil, context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTxsFromQueue(t *testing.T) {
	hndlr, _, cleanup := setup(t)
	defer cleanup()

	_, tx := makeTxCleanup()
	mustAddTx(t, hndlr, tx, 1)

	err := hndlr.AddTxsToQueue(nil, context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = hndlr.getTxsFromQueue(hndlr.db.NewTransaction(false), context.TODO(), 1, uint32(3000000), []*objs.Tx{tx})
	if err != nil {
		t.Fatal(err)
	}
}
