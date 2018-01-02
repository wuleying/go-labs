package block

import (
	"time"
)

// 区块结构体
type Block struct {
	Id            int64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// 创建新区块
func NewBlock(data string, prevId int64, prevBlockHash []byte) *Block {
	id := prevId + 1
	block := &Block{id, time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 创建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", 0, []byte{})
}
