package server

import (
	"fmt"
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/utils"
	"net"
)

const protocol = "tcp"
const nodeVersion = 1
const commandLength = 12

var nodeAddress string
var miningAddress string
var knowNodes = []string{"localhost:13000"}
var blocksInTransit = [][]byte{}
var mempool = make(map[string]b.Transaction)

type addr struct {
	AddressList []string
}

type block struct {
	AddressFrom string
	Block       []byte
}

type getBlocks struct {
	AddressFrom string
}

type getData struct {
	AddressFrom string
	Type        string
	Id          []byte
}

type inv struct {
	AddressFrom string
	Type        string
	Items       [][]byte
}

type tx struct {
	AddressFrom string
	Transaction []byte
}

type version struct {
	Version     int
	BestHeight  int
	AddressFrom string
}

func commandToBytes(command string) []byte {
	var bytes [commandLength]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func extractCommand(request []byte) []byte {
	return request[:commandLength]
}

func requestBlocks() {
	for _, node := range knowNodes {
		sendGetBlocks(node)
	}
}

func sendAddress(address string) {
	nodes := addr{knowNodes}
	nodes.AddressList = append(nodes.AddressList, nodeAddress)
	payload := utils.GobEncode(nodes)
	request := append(commandToBytes("address"), payload...)

	sendData(address, request)
}

func sendBlock(address string, b *b.Block) {
	data := block{nodeAddress, b.Serialize()}
	payload := utils.GobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(address, request)
}

func sendInv(address string, kind string, items [][]byte) {
	inventory := inv{nodeAddress, kind, items}

	payload := utils.GobEncode(inventory)
	request := append(commandToBytes("inv"), payload...)

	sendData(address, request)
}

func sendGetBlocks(address string) {
	payload := utils.GobEncode(getBlocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(address, request)
}

func sendGetData(address string, kind string, id []byte) {
	payload := utils.GobEncode(getData{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func sendTx(address string, transaction *b.Transaction) {
	data := tx{nodeAddress, transaction.Serialize()}
	payload := utils.GobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(address, request)
}

func sendVersion(address string, bc *b.BlockChain) {
	bestHeight := bc.GetBestHeight()
	payload := utils.GobEncode(version{nodeVersion, bestHeight, nodeAddress})
	request := append(commandToBytes("version"), payload...)

	sendData(address, request)
}

func sendData(address string, data []byte) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		clog.Info("% is not vaailable", address)

		var updateNodes []string

		for _, node := range knowNodes {
			if node != address {
				updateNodes = append(updateNodes, node)
			}
		}

		knowNodes = updateNodes

		return
	}

	defer conn.Close()
}
