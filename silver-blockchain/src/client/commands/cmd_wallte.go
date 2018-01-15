package commands

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/utils"
	"go-labs/silver-blockchain/src/wallet"
	"log"
)

// 创建钱包
func createWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("You new address: %s\n", address)
}

// 获取钱包地址列表
func getWalletAddresses() {
	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Printf("Address: %s, Balance: %d\n", address, balance(address))
	}
}

// 获取钱包余额
func balance(address string) int {
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

	return balance
}
