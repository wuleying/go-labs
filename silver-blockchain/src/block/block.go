package block

import (
	"bytes"
	"encoding/gob"
	"log"
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

// 序列化区块对象
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化区块对象
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block
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
