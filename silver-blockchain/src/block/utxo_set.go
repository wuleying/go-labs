package block

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
)

const UTXOBucket = "chainstate"

type UTXOSet struct {
	BlockChain *BlockChain
}

func (u UTXOSet) FindSpendableOutputs(publicKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	accmulated := 0
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UTXOBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tId := hex.EncodeToString(k)
			outs := DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.IsLockedWithKey(publicKeyHash) && accmulated < amount {
					accmulated += out.Value
					unspentOutputs[tId] = append(unspentOutputs[tId], outIdx)
				}
			}
		}

		return nil
	})

	if err != nil {
		clog.Error(0, err.Error())
	}

	return accmulated, unspentOutputs
}

func (u UTXOSet) Update(block *Block) {
	db := u.BlockChain.Db

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(UTXOBucket))

		for _, tx := range block.Transactions {
			if tx.IsCoinBase() == false {
				for _, in := range tx.In {
					updateOuts := TOutputs{}
					outsBytes := b.Get(in.Id)
					outs := DeserializeOutputs(outsBytes)

					for outIdx, out := range outs.Outputs {
						if outIdx != in.Out {
							updateOuts.Outputs = append(updateOuts.Outputs, out)
						}
					}

					if len(updateOuts.Outputs) == 0 {
						err := b.Delete(in.Id)
						if err != nil {
							clog.Error(0, err.Error())
						}
					} else {
						err := b.Put(in.Id, updateOuts.Serialize())
						if err != nil {
							clog.Error(0, err.Error())
						}
					}
				}
			}

			newOutputs := TOutputs{}
			for _, out := range tx.Out {
				newOutputs.Outputs = append(newOutputs.Outputs, out)
			}

			err := b.Put(tx.Id, newOutputs.Serialize())
			if err != nil {
				clog.Error(0, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		clog.Error(0, err.Error())
	}
}
