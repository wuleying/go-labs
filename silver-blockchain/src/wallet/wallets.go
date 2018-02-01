package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/go-labs/silver-blockchain/src/util"
	"io/ioutil"
	"os"
)

var walletFile = fmt.Sprintf("%s/%s", util.ROOT_DIR, "db/wallet.dat")

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()

	return &wallets, err
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws *Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	var wallets Wallets

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))

	err = decoder.Decode(&wallets)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	ws.Wallets = wallets.Wallets

	return nil
}

func (ws *Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)

	err := encoder.Encode(ws)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		clog.Fatal(util.CLOG_SKIP_DISPLAY_INFO, err.Error())
	}
}
