package main

import (
	"flag"
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *b.Blockchain
}

// 打印命令行使用说明
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("    add -d [BLOCK_DATA] \t Add a block to the blockchain.")
	fmt.Println("    get -h [BLOCK_HASH] \t Get a block inf by hash.")
	fmt.Println("    print \t\t\t Print all the blocks of the blockchain.")
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
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)

	pow := b.NewProofOfWork(block)
	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
}

// 添加区块
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Add block success!")
}

// 根据hash值获取区块信息
func (cli *CLI) getBlock(hashKey string) {
	block := cli.bc.GetBlock(hashKey)

	cli.printBlockInfo(block)
}

// 打印全部区块链数据
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		cli.printBlockInfo(block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// 运行命令行
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getBlockCmd := flag.NewFlagSet("get", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	addBlockData := addBlockCmd.String("d", "", "Block data")
	BlockHashKey := getBlockCmd.String("h", "", "Block hash")

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

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			cli.printUsage()
			os.Exit(1)
		}

		cli.addBlock(*addBlockData)
	}

	if getBlockCmd.Parsed() {
		if *BlockHashKey == "" {
			cli.printUsage()
			os.Exit(1)
		}

		cli.getBlock(*BlockHashKey)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
