package client

import (
	"github.com/urfave/cli"
	"go-labs/silver-blockchain/src/client/commands"
	"os"
)

/*
 * e.g.
 * ./silver-blockchain balance get -address=ADDRESS                                 Get balance info of address
 * ./silver-blockchain blockchain create -address=ADDRESS                           Create a blockchain and send genesis block reward to address
 * ./silver-blockchain blockchain print                                             Print all blockchain info
 * ./silver-blockchain transaction send -from=ADDRESS -to=ADDRESS -amount=AMOUNT    Send AMOUNT of coins from FROM address A to address B
 */

// 运行命令行
func Run() {
	app := cli.NewApp()
	app.Name = "Silver"
	app.Usage = "Blockchain client tools"
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
