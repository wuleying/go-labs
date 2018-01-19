package client

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/urfave/cli"
	"go-labs/silver-blockchain/src/client/commands"
	"os"
)

func init() {
	if err := clog.New(clog.CONSOLE, clog.ConsoleConfig{
		Level:      clog.INFO,
		BufferSize: 100},
	); err != nil {
		fmt.Printf("init console log failed. error %+v.", err)
		os.Exit(1)
	}
}

/*
 * e.g.
 *
 * ./silver-blockchain blockchain create -address=ADDRESS                           Create a blockchain and send genesis block reward to address
 * ./silver-blockchain blockchain get                                               Get all blockchain info
 * ./silver-blockchain transaction send -from=ADDRESS -to=ADDRESS -amount=AMOUNT    Send AMOUNT of coins from address A to address B
 * ./silver-blockchain wallet create                                                Create a new wallet
 * ./silver-blockchain wallet get                                                   Get all wallets address
 * ./silver-blockchain wallet balance -address=ADDRESS                              Get wallet balance info of address
 */
// 运行命令行
func Run() {
	defer clog.Shutdown()

	app := cli.NewApp()
	app.Name = "Silver Blockchain"
	app.Usage = "Client tools"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Liang Luo",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commands.Commands
	app.Run(os.Args)
}
