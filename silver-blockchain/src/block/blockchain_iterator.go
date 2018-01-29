package block

import (
	"github.com/boltdb/bolt"
	"github.com/go-clog/clog"
	"go-labs/silver-blockchain/src/util"
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
		b := tx.Bucket([]byte(util.BLOCK_BUCKET_NAME))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	i.currentHash = block.PrevBlockHash

	return block
}
