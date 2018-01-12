package client

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/wallet"
	"log"
	"strconv"
)

// 创建区块链
func (cli *CLI) createBlockChain(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("Error: Address is not valid.")
	}

	bc := b.CreateBlockChain(address)
	bc.Db.Close()
	fmt.Println("Create blockchain success.")
}

// 打印全部区块链数据
func (cli *CLI) printBlockChain() {
	bc := b.NewBlockChain("")
	defer bc.Db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		cli.printBlockInfo(block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 打印区块信息
func (cli *CLI) printBlockInfo(block *b.Block) {
	fmt.Printf("Id: #%d\n", block.Id)
	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("Hash: %x\n", block.Hash)

	pow := b.NewProofOfWork(block)
	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
}
