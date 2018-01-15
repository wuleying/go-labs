package commands

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
)

// 交易货币
func sendCoin(from string, to string, amount int) {
	bc := b.NewBlockChain(from)
	defer bc.Db.Close()

	t := b.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*b.Transaction{t})
	fmt.Println("Send success!")
}
