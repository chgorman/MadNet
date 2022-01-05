package txfees

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/MadBase/MadNet/application/objs/uint256"
	"github.com/MadBase/MadNet/crypto"
)

func TestTxHeap(t *testing.T) {
	txhash1 := crypto.Hasher([]byte("TxHash1"))
	utxoID11 := crypto.Hasher([]byte("utxoID11"))
	utxoID12 := crypto.Hasher([]byte("utxoID12"))
	value1, err := new(uint256.Uint256).FromUint64(1)
	if err != nil {
		t.Fatal(err)
	}
	item1 := &TxItem{
		txhash:  txhash1,
		value:   value1,
		utxoIDs: [][]byte{utxoID11, utxoID12},
	}

	txhash2 := crypto.Hasher([]byte("TxHash2"))
	utxoID21 := crypto.Hasher([]byte("utxoID21"))
	utxoID22 := crypto.Hasher([]byte("utxoID22"))
	value2, err := new(uint256.Uint256).FromUint64(5)
	if err != nil {
		t.Fatal(err)
	}
	item2 := &TxItem{
		txhash:  txhash2,
		value:   value2,
		utxoIDs: [][]byte{utxoID21, utxoID22},
	}

	txhash3 := crypto.Hasher([]byte("TxHash3"))
	utxoID31 := crypto.Hasher([]byte("utxoID31"))
	utxoID32 := crypto.Hasher([]byte("utxoID32"))
	value3, err := new(uint256.Uint256).FromUint64(13)
	if err != nil {
		t.Fatal(err)
	}
	item3 := &TxItem{
		txhash:  txhash3,
		value:   value3,
		utxoIDs: [][]byte{utxoID31, utxoID32},
	}

	txhash4 := crypto.Hasher([]byte("TxHash4"))
	utxoID41 := crypto.Hasher([]byte("utxoID41"))
	utxoID42 := crypto.Hasher([]byte("utxoID42"))
	value4, err := new(uint256.Uint256).FromUint64(7)
	if err != nil {
		t.Fatal(err)
	}
	item4 := &TxItem{
		txhash:  txhash4,
		value:   value4,
		utxoIDs: [][]byte{utxoID41, utxoID42},
	}

	txhash5 := crypto.Hasher([]byte("TxHash5"))
	utxoID51 := crypto.Hasher([]byte("utxoID51"))
	utxoID52 := crypto.Hasher([]byte("utxoID52"))
	value5, err := new(uint256.Uint256).FromUint64(11)
	if err != nil {
		t.Fatal(err)
	}
	item5 := &TxItem{
		txhash:  txhash5,
		value:   value5,
		utxoIDs: [][]byte{utxoID51, utxoID52},
	}

	items := []*TxItem{}
	items = append(items, item1)
	items = append(items, item2)
	items = append(items, item3)
	items = append(items, item4)
	items = append(items, item5)

	txh := make(TxHeap, 0, len(items))
	heap.Init(&txh)
	for k := 0; k < len(items); k++ {
		heap.Push(&txh, items[k])
	}
	fmt.Println("Popping from heap:")
	for txh.Len() > 0 {
		item := heap.Pop(&txh).(*TxItem)
		fmt.Printf("Hash: %x; Value: %v\n", item.txhash, item.value)
	}
}
