package commands

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/wallet"
	"log"
)

// 交易货币
func sendCoin(from string, to string, amount int) {
	if !wallet.ValidateAddress(from) {
		log.Panic("Error:Sender address is not valid.")
	}

	if !wallet.ValidateAddress(to) {
		log.Panic("Error:Recipient address is not valid.")
	}

	bc := b.NewBlockChain(from)
	defer bc.Db.Close()

	t := b.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*b.Transaction{t})
	fmt.Println("Send success!")
}
