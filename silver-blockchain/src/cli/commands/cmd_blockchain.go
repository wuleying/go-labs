package commands

import (
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/util"
	"go-labs/silver-blockchain/src/wallet"
	"strconv"
)

// 创建区块链
func createBlockChain(address string, nodeId string) {
	if !wallet.ValidateAddress(address) {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, "Address [%s] is not valid.", address)
	}

	bc := b.CreateBlockChain(address, nodeId)
	defer bc.Db.Close()

	UTXOSet := b.UTXOSet{bc}
	UTXOSet.Reindex()

	clog.Info("Create blockchain success.")
}

// 打印全部区块链数据
func getBlockChain(nodeId string) {
	bc := b.NewBlockChain(nodeId)
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
	clog.Info("############################## Block #%d ##############################", block.Id)
	clog.Info("PrevBlockHash: %x", block.PrevBlockHash)
	clog.Info("Hash: %x", block.Hash)
	pow := b.NewProofOfWork(block)
	clog.Info("Pow: %s\n", strconv.FormatBool(pow.Validate()))
}
