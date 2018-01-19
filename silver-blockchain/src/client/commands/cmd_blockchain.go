package commands

import (
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/wallet"
	"strconv"
)

// 创建区块链
func createBlockChain(address string) {
	if !wallet.ValidateAddress(address) {
		clog.Fatal(2, "Address [%s] is not valid.", address)
	}

	bc := b.CreateBlockChain(address)
	bc.Db.Close()

	clog.Info("Create blockchain success.")
}

// 打印全部区块链数据
func GetBlockChain() {
	bc := b.NewBlockChain("")
	defer bc.Db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		formatBlockInfo(block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 格式化输出区块信息
func formatBlockInfo(block *b.Block) {
	clog.Info("Id: #%d\n", block.Id)
	clog.Info("PrevBlockHash: %x\n", block.PrevBlockHash)
	clog.Info("Hash: %x\n", block.Hash)
	pow := b.NewProofOfWork(block)
	clog.Info("Pow: %s\n", strconv.FormatBool(pow.Validate()))
}
