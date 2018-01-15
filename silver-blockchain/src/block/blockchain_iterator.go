package block

import (
	"github.com/boltdb/bolt"
	"log"
)

// 区块链迭代器结构体
type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// 区块链迭代器
func (bc *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{bc.Tip, bc.Db}

	return bci
}

// 迭代区获取区块信息
func (i *BlockChainIterator) Next() *Block {
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
