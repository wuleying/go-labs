package server

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/go-clog/clog"
	b "go-labs/silver-blockchain/src/block"
	"go-labs/silver-blockchain/src/util"
	"io/ioutil"
	"net"
)

const (
	PROTOCOL       = "tcp"
	NODE_VERSION   = 1
	COMMAND_LENGTH = 12
)

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
	var bytes [COMMAND_LENGTH]byte

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
	return request[:COMMAND_LENGTH]
}

func requestBlocks() {
	for _, node := range knowNodes {
		sendGetBlocks(node)
	}
}

func sendAddress(address string) {
	nodes := addr{knowNodes}
	nodes.AddressList = append(nodes.AddressList, nodeAddress)
	payload := util.GobEncode(nodes)
	request := append(commandToBytes("address"), payload...)

	sendData(address, request)
}

func sendBlock(address string, b *b.Block) {
	data := block{nodeAddress, b.Serialize()}
	payload := util.GobEncode(data)
	request := append(commandToBytes("block"), payload...)

	sendData(address, request)
}

func sendInv(address string, kind string, items [][]byte) {
	inventory := inv{nodeAddress, kind, items}

	payload := util.GobEncode(inventory)
	request := append(commandToBytes("inv"), payload...)

	sendData(address, request)
}

func sendGetBlocks(address string) {
	payload := util.GobEncode(getBlocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(address, request)
}

func sendGetData(address string, kind string, id []byte) {
	payload := util.GobEncode(getData{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func sendTx(address string, transaction *b.Transaction) {
	data := tx{nodeAddress, transaction.Serialize()}
	payload := util.GobEncode(data)
	request := append(commandToBytes("tx"), payload...)

	sendData(address, request)
}

func sendVersion(address string, bc *b.BlockChain) {
	bestHeight := bc.GetBestHeight()
	payload := util.GobEncode(version{NODE_VERSION, bestHeight, nodeAddress})
	request := append(commandToBytes("version"), payload...)

	sendData(address, request)
}

func sendData(address string, data []byte) {
	conn, err := net.Dial(PROTOCOL, address)
	if err != nil {
		clog.Info("%s is not available", address)

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

func handleAddress(request []byte) {
	var buff bytes.Buffer
	var payload addr

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	knowNodes = append(knowNodes, payload.AddressList...)
	clog.Info("There are %d known nodes now!", len(knowNodes))
	requestBlocks()
}

func handleBlock(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	blockData := payload.Block
	blockInfo := b.DeserializeBlock(blockData)
	clog.Info("Recevied a new block!")

	bc.AddBlock(blockInfo)
	clog.Info("Added block %x", blockInfo.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddressFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	} else {
		UTXOSet := b.UTXOSet{bc}
		UTXOSet.Reindex()
	}
}

func handleInv(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	clog.Info("Recevied inventory with %d %s", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		sendGetData(payload.AddressFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if bytes.Compare(b, blockHash) != 0 {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}

	if payload.Type == "tx" {
		tId := payload.Items[0]

		if mempool[hex.EncodeToString(tId)].Id == nil {
			sendGetData(payload.AddressFrom, "tx", tId)
		}
	}
}

func handleGetBlocks(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload getBlocks

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	blocks := bc.GetBlockHashes()
	sendInv(payload.AddressFrom, "block", blocks)
}

func handleGetData(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload getData

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	if payload.Type == "block" {
		block, err := bc.GetBlock([]byte(payload.Id))
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		sendBlock(payload.AddressFrom, &block)
	}

	if payload.Type == "tx" {
		tId := hex.EncodeToString(payload.Id)
		tx := mempool[tId]

		sendTx(payload.AddressFrom, &tx)
	}
}

func handleTx(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload tx

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	txData := payload.Transaction
	tx := b.DeserializeTransaction(txData)
	mempool[hex.EncodeToString(tx.Id)] = tx

	if nodeAddress == knowNodes[0] {
		for _, node := range knowNodes {
			if node != nodeAddress && node != payload.AddressFrom {
				sendInv(node, "tx", [][]byte{tx.Id})
			}
		}
	} else {
		if len(mempool) >= 2 && len(miningAddress) > 0 {
		MineTransactions:
			var txs []*b.Transaction

			for id := range mempool {
				tx := mempool[id]
				if bc.VerifyTransaction(&tx) {
					txs = append(txs, &tx)
				}
			}

			if len(txs) == 0 {
				clog.Info("All transactions are invalid! Waiting for new ones.")
				return
			}

			cbTx := b.NewCoinBase(miningAddress, "")
			txs = append(txs, cbTx)

			newBlock := bc.MineBlock(txs)
			UTXOSet := b.UTXOSet{bc}
			UTXOSet.Reindex()
			clog.Info("New block is mined.")

			for _, tx := range txs {
				txId := hex.EncodeToString(tx.Id)
				delete(mempool, txId)
			}

			for _, node := range knowNodes {
				if node != nodeAddress {
					sendInv(node, "block", [][]byte{newBlock.Hash})
				}
			}

			if len(mempool) > 0 {
				goto MineTransactions
			}
		}
	}
}

func handleVersion(request []byte, bc *b.BlockChain) {
	var buff bytes.Buffer
	var payload version

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if myBestHeight < foreignerBestHeight {
		sendGetBlocks(payload.AddressFrom)
	} else if myBestHeight > foreignerBestHeight {
		sendVersion(payload.AddressFrom, bc)
	}

	if !nodeIsKnown(payload.AddressFrom) {
		knowNodes = append(knowNodes, payload.AddressFrom)
	}
}

func handleConnection(conn net.Conn, bc *b.BlockChain) {
	request, err := ioutil.ReadAll(conn)
	if err != nil {
		clog.Fatal(2, err.Error())
	}

	command := bytesToCommand(request[:COMMAND_LENGTH])
	clog.Info("Received %s command", command)

	switch command {
	case "address":
		handleAddress(request)
	case "block":
		handleBlock(request, bc)
	case "inv":
		handleInv(request, bc)
	case "getblocks":
		handleGetBlocks(request, bc)
	case "getdata":
		handleGetData(request, bc)
	case "tx":
		handleTx(request, bc)
	case "version":
		handleVersion(request, bc)
	default:
		clog.Info("Unknown command.")
	}

	conn.Close()
}

func nodeIsKnown(address string) bool {
	for _, node := range knowNodes {
		if node == address {
			return true
		}
	}

	return false
}

func StartServer(nodeId string, minerAddress string) {
	nodeAddress = fmt.Sprintf("localhost:%s", nodeId)
	miningAddress = minerAddress
	ln, err := net.Listen(PROTOCOL, nodeAddress)
	if err != nil {
		clog.Fatal(2, err.Error())
	}
	defer ln.Close()

	bc := b.NewBlockChain(nodeId)

	if nodeAddress != knowNodes[0] {
		sendVersion(knowNodes[0], bc)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			clog.Fatal(2, err.Error())
		}

		go handleConnection(conn, bc)
	}
}
