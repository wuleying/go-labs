package block

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
	"github.com/juju/errors"
	"os"
)

// 数据库路径
const dbFileTemp = "db/silver-blockchain-%s.db"

// Last hash key
const lastHashKey = "lastHash"

// Bucket名称
const blockBucket = "blocks"

// 创世币数据
const genesisCoinBaseData = "hello luoliang"

// 区块链结构体
type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

// 创建区块链
func CreateBlockChain(address string, nodeId string) *BlockChain {
	dbFile := fmt.Sprintf(dbFileTemp, nodeId)
	if dbExists(dbFile) {
		clog.Fatal(2, "Blockchain already exists.")
	}

	var tip []byte

	coinBase := NewCoinBase(address, genesisCoinBaseData)
	genesisBlock := NewGenesisBlock(coinBase)

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		err = b.Put([]byte(lastHashKey), genesisBlock.Hash)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		tip = genesisBlock.Hash

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	bc := BlockChain{tip, db}

	return &bc
}

// 创建新区块链
func NewBlockChain(nodeId string) *BlockChain {
	dbFile := fmt.Sprintf(dbFileTemp, nodeId)
	if dbExists(dbFile) == false {
		clog.Fatal(2, "No existing blockchain found. Create one first.")
	}

	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		tip = b.Get([]byte(lastHashKey))

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	bc := BlockChain{tip, db}

	return &bc
}

func (bc *BlockChain) GetBlockHashes() [][]byte {
	var blocks [][]byte
	bci := bc.Iterator()

	for {
		block := bci.Next()

		blocks = append(blocks, block.Hash)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return blocks
}

// 挖矿
func (bc *BlockChain) MineBlock(transactions []*Transaction) *Block {
	var lastHash []byte
	var lastBlock *Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte(lastHashKey))
		lastBlock = DeserializeBlock(b.Get(lastHash))

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	newBlock := NewBlock(transactions, lastBlock.Id, lastBlock.Hash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		err = b.Put([]byte(lastHashKey), newBlock.Hash)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		bc.Tip = newBlock.Hash

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return newBlock
}

// 添加区块
func (bc *BlockChain) AddBlock(transactions []*Transaction) {
	var lastHash []byte
	var lastBlock *Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte(lastHashKey))
		lastBlock = DeserializeBlock(b.Get(lastHash))

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	newBlock := NewBlock(transactions, lastBlock.Id, lastBlock.Hash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			clog.Fatal(2, err.Error())
		}

		err = b.Put([]byte(lastHashKey), newBlock.Hash)

		if err != nil {
			clog.Fatal(2, err.Error())
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

// 根据hash获取区块信息
func (bc *BlockChain) GetBlock(blockHash []byte) (Block, error) {
	var block Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		blockData := b.Get(blockHash)
		if blockData == nil {
			return errors.New("Block is not found.")
		}

		block = *DeserializeBlock(blockData)
		return nil
	})

	if err != nil {
		return block, err
	}

	return block, nil
}

func (bc *BlockChain) FindUTXO() map[string]TOutputs {
	UTXO := make(map[string]TOutputs)
	spentUTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, t := range block.Transactions {
			tId := hex.EncodeToString(t.Id)

		Outputs:
			for outIdx, out := range t.Out {
				if spentUTXOs[tId] != nil {
					for _, spentOutIdx := range spentUTXOs[tId] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				outs := UTXO[tId]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[tId] = outs
			}

			if t.IsCoinBase() == false {
				for _, in := range t.In {
					inId := hex.EncodeToString(in.Id)
					spentUTXOs[inId] = append(spentUTXOs[inId], in.Out)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

func (bc *BlockChain) FindSpendableOutputs(publicKeyHash []byte, amount int) (int, map[string][]int) {
	accountMulated := 0
	unspentOutputs := make(map[string][]int)
	unspentT := bc.FindUnspentTransactions(publicKeyHash)

Work:
	for _, t := range unspentT {
		tId := hex.EncodeToString(t.Id)

		for outIdx, out := range t.Out {
			if out.IsLockedWithKey(publicKeyHash) && accountMulated < amount {
				accountMulated += out.Value

				unspentOutputs[tId] = append(unspentOutputs[tId], outIdx)

				if accountMulated >= amount {
					break Work
				}
			}
		}
	}

	return accountMulated, unspentOutputs
}

func (bc *BlockChain) FindTransaction(id []byte) (Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, t := range block.Transactions {
			if bytes.Compare(t.Id, id) == 0 {
				return *t, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
}

func (bc *BlockChain) FindUnspentTransactions(publicKeyHash []byte) []Transaction {
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

				if out.IsLockedWithKey(publicKeyHash) {
					unspentT = append(unspentT, *t)
				}
			}

			if t.IsCoinBase() == false {
				for _, in := range t.In {
					if in.UsesKey(publicKeyHash) {
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

func (bc *BlockChain) SignTransaction(t *Transaction, privateKey ecdsa.PrivateKey) {
	prevTs := make(map[string]Transaction)

	for _, in := range t.In {
		prevT, err := bc.FindTransaction(in.Id)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		prevTs[hex.EncodeToString(prevT.Id)] = prevT
	}

	t.Sign(privateKey, prevTs)
}

func (bc *BlockChain) VerifyTransaction(t *Transaction) bool {
	prevTs := make(map[string]Transaction)

	for _, in := range t.In {
		prevT, err := bc.FindTransaction(in.Id)
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		prevTs[hex.EncodeToString(prevT.Id)] = prevT
	}

	return t.Verify(prevTs)
}

// 区块链迭代器
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.Tip, bc.Db}

	return bci
}

func (bc *BlockChain) GetBestHeight() int {
	var lastBlock Block

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash := b.Get([]byte(lastHashKey))
		blockData := b.Get(lastHash)
		lastBlock = *DeserializeBlock(blockData)

		return nil
	})

	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return lastBlock.Height
}

// 检查数据库文件是否存在
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
