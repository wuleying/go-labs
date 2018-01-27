package commands

import (
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/util"
	"go-labs/silver-blockchain/src/wallet"
)

// 创建钱包
func createWallet() {
	wallets, _ := wallet.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	clog.Info("You new address: %s", address)
}

// 获取钱包地址列表
func getWalletAddresses(nodeId string) {
	wallets, err := wallet.NewWallets()
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		clog.Info("Address: %s, Balance: %d", address, balance(address, nodeId))
	}
}

// 获取钱包余额
func balance(address string, nodeId string) int {
	if !wallet.ValidateAddress(address) {
		clog.Fatal(2, "Address %s is not valid.", address)
	}

	bc := b.NewBlockChain(nodeId)
	defer bc.Db.Close()

	UTXOSet := b.UTXOSet{bc}

	balance := 0
	publicKeyHash := util.Base58Decode([]byte(address))
	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-util.ADDRESS_CHECKSUM_LEN]

	UTXOs := UTXOSet.FindUTXO(publicKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}
