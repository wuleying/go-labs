package block

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/juju/errors"
	"log"
	"os"
)

// 数据库路径
const dbFile = "db/silver-blockchain.db"

// Last hash key
const lastHashKey = "lastHash"

// Bucket名称
const blockBucket = "blocks"

// 创世币数据
const genesisCoinbaseData = "hello luoliang"

// 区块链结构体
type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

// 区块链迭代器结构体
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// 挖矿
func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lastHash []byte
	var lastBlock *Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte(lastHashKey))
		lastBlock = DeserializeBlock(b.Get(lastHash))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastBlock.Id, lastBlock.Hash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte(lastHashKey), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

// 添加区块
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	var lastHash []byte
	var lastBlock *Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte(lastHashKey))
		lastBlock = DeserializeBlock(b.Get(lastHash))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastBlock.Id, lastBlock.Hash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte(lastHashKey), newBlock.Hash)

		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

// 根据id获取区块信息
func (bc *Blockchain) GetBlock(id int64) (*Block, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		if block.Id == id {
			return block, nil
		}
	}

	return nil, errors.New("Block is not exist.")
}

// 区块链迭代器
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.Db}

	return bci
}

// 迭代区获取区块信息
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

func (bc *Blockchain) FindUTXO(address string) []TOutput {
	var UTXO []TOutput

	unspentTransactions := bc.FineUnspentTransactions(address)

	for _, t := range unspentTransactions {
		for _, out := range t.Out {
			if out.CanBeUnlockWith(address) {
				UTXO = append(UTXO, out)
			}
		}
	}

	return UTXO
}

func (bc *Blockchain) FineUnspentTransactions(address string) []Transaction {
	var unspentT []Transaction
	spentTXO := make(map[string][]int)

	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, t := range block.Transactions {
			tId := hex.EncodeToString(t.Id)

		Outputs:
			for outId, out := range t.Out {
				if spentTXO[tId] != nil {
					for _, spentOut := range spentTXO[tId] {
						if spentOut == outId {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockWith(address) {
					unspentT = append(unspentT, *t)
				}
			}

			if t.IsCoinbase() == false {
				for _, in := range t.In {
					if in.CanUnlockOutputWith(address) {
						inTId := hex.EncodeToString(in.Id)
						spentTXO[inTId] = append(spentTXO[inTId], in.Out)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentT
}

// 检查数据库文件是否存在
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

// 创建新区块链
func NewBlockchain(address string) *Blockchain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		tip = b.Get([]byte(lastHashKey))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

// 创建区块链
func CreateBlockchain(address string) *Blockchain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbt := NewCoinbase(address, genesisCoinbaseData)
		genesisBlock := NewGenesisBlock(cbt)

		b, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte(lastHashKey), genesisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		tip = genesisBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}
