package block

import (
	"encoding/hex"
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/util"
)

type UTXOSet struct {
	BlockChain *BlockChain
}

func (u UTXOSet) FindSpendableOutputs(publicKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	accumulated := 0
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(util.UTXO_BUCKET_NAME)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tId := hex.EncodeToString(k)
			outs := DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.IsLockedWithKey(publicKeyHash) && accumulated < amount {
					accumulated += out.Value
					unspentOutputs[tId] = append(unspentOutputs[tId], outIdx)
				}
			}
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return accumulated, unspentOutputs
}

func (u UTXOSet) FindUTXO(publicKeyHash []byte) []TOutput {
	var UTXOs []TOutput
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(util.UTXO_BUCKET_NAME)).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			outs := DeserializeOutputs(v)

			for _, out := range outs.Outputs {
				if out.IsLockedWithKey(publicKeyHash) {
					UTXOs = append(UTXOs, out)
				}
			}
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return UTXOs
}

func (u UTXOSet) CountTransactions() int {
	counter := 0
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(util.UTXO_BUCKET_NAME)).Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			counter++
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	return counter
}

func (u UTXOSet) Reindex() {
	db := u.BlockChain.Db
	bucketName := []byte(util.UTXO_BUCKET_NAME)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil && err != bolt.ErrBucketNotFound {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	UTXO := u.BlockChain.FindUTXO()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		for tId, outs := range UTXO {
			key, err := hex.DecodeString(tId)
			if err != nil {
				clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
			}

			err = b.Put(key, outs.Serialize())
			if err != nil {
				clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
}

func (u UTXOSet) Update(block *Block) {
	db := u.BlockChain.Db

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(util.UTXO_BUCKET_NAME))

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
							clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
						}
					} else {
						err := b.Put(in.Id, updateOuts.Serialize())
						if err != nil {
							clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
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
				clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
}
