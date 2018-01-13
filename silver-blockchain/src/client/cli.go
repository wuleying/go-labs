package client

import (
	"github.com/urfave/cli"
	"go-labs/silver-blockchain/src/client/commands"
	"os"
)

// 运行命令行
func Run() {
	app := cli.NewApp()
	app.Name = "Silver blockchain"
	app.Usage = "client tools"
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
