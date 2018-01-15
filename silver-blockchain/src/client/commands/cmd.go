package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:    "balance",
		Aliases: []string{"b"},
		Usage:   "operation",
		Subcommands: []cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get balance info of address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "address",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.String("address")) < 1 {
						return cli.ShowAppHelp(c)
					}

					getBalance(c.String("address"))
					return nil
				},
			},
		},
	},
	{
		Name:    "blockchain",
		Aliases: []string{"bc"},
		Usage:   "operation",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a blockchain and send genesis block reward to address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "address",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.String("address")) < 1 {
						return cli.ShowAppHelp(c)
					}

					createBlockChain(c.String("address"))
					return nil
				},
			},
			{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "Print all blockchain info",
				Action: func(c *cli.Context) error {
					printBlockChain()
					return nil
				},
			},
		},
	},
}
