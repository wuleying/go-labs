package commands

import (
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/util"
	"go-labs/silver-blockchain/src/wallet"
)

// 交易货币
func sendCoin(from string, to string, amount int, nodeId string, mineNow bool) {
	if !wallet.ValidateAddress(from) {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, "Sender address [%s] is not valid.", from)
	}

	if !wallet.ValidateAddress(to) {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, "Recipient address is not valid.", to)
	}

	bc := b.NewBlockChain(nodeId)
	UTXOSet := b.UTXOSet{bc}
	defer bc.Db.Close()

	t := b.NewUTXOTransaction(from, to, amount, &UTXOSet)
	coinBaseT := b.NewCoinBase(from, "")
	ts := []*b.Transaction{coinBaseT, t}

	newBlock := bc.MineBlock(ts)
	UTXOSet.Update(newBlock)

	clog.Info("Send success!")
}
