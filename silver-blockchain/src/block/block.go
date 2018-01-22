package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/go-clog/clog"
	"time"
)

// 区块结构体
type Block struct {
	Id            int64
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

// 创建新区块
func NewBlock(transactions []*Transaction, prevId int64, prevBlockHash []byte) *Block {
	id := prevId + 1
	block := &Block{id, time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, 1}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 创建创世区块
func NewGenesisBlock(coinBase *Transaction) *Block {
	return NewBlock([]*Transaction{coinBase}, 0, []byte{})
}

// 交易哈希
func (b *Block) HashTransaction() []byte {
	var tHashes [][]byte
	var tHash [32]byte

	for _, t := range b.Transactions {
		tHashes = append(tHashes, t.Id)
	}
	tHash = sha256.Sum256(bytes.Join(tHashes, []byte{}))

	return tHash[:]
}

// 序列化区块对象
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return result.Bytes()
}

// 反序列化区块对象
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	return &block
}
