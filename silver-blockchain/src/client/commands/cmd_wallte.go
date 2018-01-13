package commands

import (
	"fmt"
	"go-labs/silver-blockchain/src/wallet"
	"log"
)

// 创建钱包
func createWallet() {
	wallets, err := wallet.NewWallets()
	if err != nil {
		log.Panic(err)
	}

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
		fmt.Println(address)
	}
}
