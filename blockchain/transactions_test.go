package blockchain

import (
	"testing"

	"github.com/hunman89/nomadcoin/utils"
)

const address string = "d58f94e39917a976422a0f362a582aa21ca22cdcad400dea0f0a21f340c42b3f6e36680722542bedf7a03c2b4a05d87e9dec2abd7bc4d77d00eec45102d0252b"

func TestMakeTx(t *testing.T) {
	// t.Run("Tx not make", func(t *testing.T) {
	// 	dbStorage = fakeDB{
	// 		fakeLoadChain: func() []byte {
	// 			bc := &blockchain{
	// 				Height:            2,
	// 				NewestHash:        "xxx",
	// 				CurrentDifficulty: 1,
	// 			}
	// 			return utils.ToBytes(bc)
	// 		},
	// 		fakeFindBlock: func() []byte {
	// 			b := &Block{
	// 				Height: 2,
	// 				Transactions: []*Tx{
	// 					{ID: "test", TxIns: []*TxIn{{TxID: "aa", Index: 0, Signature: "COINBASE"}}, TxOuts: []*TxOut{}},
	// 				},
	// 			}
	// 			return utils.ToBytes(b)
	// 		},
	// 	}
	// 	tx, _ := makeTx("xx", "yy", 20)
	// 	if tx != nil {
	// 		t.Error("tx shoud not make")
	// 	}
	// })
	t.Run("Tx maked", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				bc := &blockchain{
					Height:            2,
					NewestHash:        "xxx",
					CurrentDifficulty: 1,
				}
				return utils.ToBytes(bc)
			},
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 2,
					Transactions: []*Tx{
						{ID: "test", TxIns: []*TxIn{{TxID: "aa", Index: 0, Signature: "COINBASE"}}, TxOuts: []*TxOut{{Address: address, Amount: 50}}},
					},
				}
				return utils.ToBytes(b)
			},
		}
		tx, _ := makeTx(address, "yy", 20)
		if tx == nil {
			t.Error("tx shoud make")
		}
	})
}
