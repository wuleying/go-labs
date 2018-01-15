package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:    "balance",
		Aliases: []string{"b"},
		Usage:   "Balance opertaions",
		Subcommands: []cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get balance info of address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "address",
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
		Usage:   "Blockchain opertaions",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a blockchain and send genesis block reward to address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "address",
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
	{
		Name:    "transaction",
		Aliases: []string{"t"},
		Usage:   "Transaction opertaions",
		Subcommands: []cli.Command{
			{
				Name:    "send",
				Aliases: []string{"s"},
				Usage:   "Send AMOUNT of coins from FROM address A to address B",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "from",
					},
					cli.StringFlag{
						Name: "to",
					},
					cli.IntFlag{
						Name: "amount",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.String("from")) < 1 || len(c.String("to")) < 1 || c.Int("amount") <= 0 {
						return cli.ShowAppHelp(c)
					}

					sendCoin(c.String("from"), c.String("to"), c.Int("amount"))
					return nil
				},
			},
		},
	},
	{
		Name:    "wallet",
		Aliases: []string{"w"},
		Usage:   "Wallet opertaions",
		Subcommands: []cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a new wallet",
				Action: func(c *cli.Context) error {
					createWallet()
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get all wallets address",
				Action: func(c *cli.Context) error {
					getWalletAddresses()
					return nil
				},
			},
		},
	},
}
