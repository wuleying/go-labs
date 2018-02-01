package cli

import (
	ucli "github.com/urfave/cli"
	"github.com/wuleying/go-labs/silver-blockchain/src/cli/commands"
	"os"
)

/*
 * e.g.
 *
 * ./silver-blockchain blockchain create -address=ADDRESS                           Create a blockchain and send genesis block reward to address
 * ./silver-blockchain blockchain get                                               Get all blockchain info
 * ./silver-blockchain transaction send -from=ADDRESS -to=ADDRESS -amount=AMOUNT    Send AMOUNT of coins from address A to address B
 * ./silver-blockchain wallet create                                                Create a new wallet
 * ./silver-blockchain wallet get                                                   Get all wallets address
 * ./silver-blockchain wallet balance -address=ADDRESS                              Get wallet balance info of address
 * ./silver-blockchain utxo reindex -address=ADDRESS                                Reindex the UTXO set
 * ./silver-blockchain node start -miner=ADDRESS                                    Start the node
 */
// 运行命令行
func Run() {
	if len(os.Getenv("NODE_ID")) < 1 {
		os.Setenv("NODE_ID", "13000")
	}

	app := ucli.NewApp()
	app.Name = "Silver Blockchain"
	app.Usage = "CLI tools"
	app.Version = "0.0.1"
	app.Authors = []ucli.Author{
		ucli.Author{
			Name:  "Luo",
			Email: "lolooo@live.com",
		},
	}
	app.Commands = commands.Commands
	app.Run(os.Args)
}
