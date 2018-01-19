package block

import (
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
)

// 区块链迭代器结构体
type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
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
		clog.Error(0, err.Error())
	}

	i.currentHash = block.PrevBlockHash

	return block
}
