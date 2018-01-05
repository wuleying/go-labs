package client

import (
	"flag"
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"log"
	"os"
	"strconv"
)

type CLI struct{}

// 打印命令行使用说明
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -d [BLOCK_DATA] \t Add a block to the blockchain.")
	fmt.Println("  get -i [BLOCK_ID] \t Get a block inf by id.")
	fmt.Println("  print \t\t Print all the blocks of the blockchain.")
}

// 校验命令行参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printBlockInfo(block *b.Block) {
	fmt.Printf("Id: #%d\n", block.Id)
	fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
	fmt.Printf("HashTransaction: %s\n", block.HashTransaction())
	fmt.Printf("Hash: %x\n", block.Hash)

	pow := b.NewProofOfWork(block)
	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
}

// 创建区块链
func (cli *CLI) createBlockchain(address string) {
	bc := b.CreateBlockchain(address)
	bc.Db.Close()
	fmt.Println("Create blockchain success.")
}

// 获取账户余额
func (cli *CLI) getBalance(address string) {
	bc := b.NewBlockchain(address)
	defer bc.Db.Close()

	balance := 0

	UTXO := bc.FindUTXO(address)

	for _, out := range UTXO {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d", address, balance)
}

func (cli *CLI) send(from string, to string, amout int) {
	bc := b.NewBlockchain(from)
	defer bc.Db.Close()

	t := b.NewUTXOTransaction(from, to, amout, bc)
	bc.MineBlock([]*b.Transaction{t})
	fmt.Println("Send success!")
}

// 运行命令行
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getBlockCmd := flag.NewFlagSet("get", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "get":
		err := getBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}
}
