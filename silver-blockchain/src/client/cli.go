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

// 打印全部区块链数据
func (cli *CLI) printAllBlockchain() {
	bc := b.NewBlockchain("")
	defer bc.Db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		cli.printBlockInfo(block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) send(from string, to string, amount int) {
	bc := b.NewBlockchain(from)
	defer bc.Db.Close()

	t := b.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*b.Transaction{t})
	fmt.Println("Send success!")
}

// 运行命令行
func (cli *CLI) Run() {
	cli.validateArgs()

	// 命令行方法
	getBalanceCmd := flag.NewFlagSet("balance", flag.ExitOnError)
	getAllCmd := flag.NewFlagSet("all", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("create", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	// 参数
	balanceAddressParam := getBalanceCmd.String("address", "", "The address to get balance for")
	blockchainAddressParam := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFromParam := sendCmd.String("from", "", "Source wallet address")
	sendToParam := sendCmd.String("to", "", "Destination wallet address")
	sendAmountParam := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "balance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "all":
		err := getAllCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "create":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}

	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *balanceAddressParam == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}

		cli.getBalance(*balanceAddressParam)
	}

	if getAllCmd.Parsed() {
		cli.printAllBlockchain()
	}

	if createBlockchainCmd.Parsed() {
		if *blockchainAddressParam == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}

		cli.createBlockchain(*blockchainAddressParam)
	}

	if sendCmd.Parsed() {
		if *sendFromParam == "" || *sendToParam == "" || *sendAmountParam <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFromParam, *sendToParam, *sendAmountParam)
	}
}