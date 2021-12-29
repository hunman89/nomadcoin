package blockchain

import (
	"testing"

	"github.com/hunman89/nomadcoin/utils"
)

func TestMakeTx(t *testing.T) {
	t.Run("Tx not make", func(t *testing.T) {
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
						{ID: "test", TxIns: []*TxIn{{TxID: "aa", Index: 0, Signature: "COINBASE"}}, TxOuts: []*TxOut{}},
					},
				}
				return utils.ToBytes(b)
			},
		}
		tx, _ := makeTx("xx", "yy", 20)
		if tx != nil {
			t.Error("tx shoud not make")
		}
	})
}
