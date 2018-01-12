package client

import (
	"flag"
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"log"
	"os"
)

// CLI结构体
type CLI struct{}

// 打印命令行使用说明
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  balance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("  blockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  all - Print all the blocks of the blockchain")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
}

// 校验命令行参数
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// 交易货币
func (cli *CLI) send(from string, to string, amount int) {
	bc := b.NewBlockChain(from)
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
	blockChainCmd := flag.NewFlagSet("blockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	// 参数
	balanceAddressParam := getBalanceCmd.String("address", "", "The address to get balance for")
	blockChainAddressParam := blockChainCmd.String("address", "", "The address to send genesis block reward to")
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

	case "blockchain":
		err := blockChainCmd.Parse(os.Args[2:])
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
		cli.printBlockChain()
	}

	if blockChainCmd.Parsed() {
		if *blockChainAddressParam == "" {
			blockChainCmd.Usage()
			os.Exit(1)
		}

		cli.createBlockChain(*blockChainAddressParam)
	}

	if sendCmd.Parsed() {
		if *sendFromParam == "" || *sendToParam == "" || *sendAmountParam <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFromParam, *sendToParam, *sendAmountParam)
	}
}
