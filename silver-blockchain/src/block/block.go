package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 区块结构体
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// 设置区块哈希
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	hash := sha256.Sum256(bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}))
	b.Hash = hash[:]
}

// 创建新区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 创建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
