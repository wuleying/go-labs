package commands

import (
	"fmt"
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/wallet"
	"strconv"
)

// 创建区块链
func createBlockChain(address string) {
	if !wallet.ValidateAddress(address) {
		clog.Error(1, "Address [%s] is not valid.", address)
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
	fmt.Printf("Id: #%d\n", block.Id)
	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("Hash: %x\n", block.Hash)

	pow := b.NewProofOfWork(block)
	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
}
