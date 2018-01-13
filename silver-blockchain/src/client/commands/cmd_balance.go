package commands

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/utils"
	"go-labs/silver-blockchain/src/wallet"
	"log"
)

// 获取账户余额
func getBalance(address string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("Error: Address is not valid.")
	}

	bc := b.NewBlockChain(address)
	defer bc.Db.Close()

	balance := 0
	publicKeyHash := utils.Base58Decode([]byte(address))
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-4]

	UTXO := bc.FindUTXO(publicKeyHash)

	for _, out := range UTXO {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", address, balance)
}
