package commands

import (
	"github.com/go-clog/clog"
	ucli "github.com/urfave/cli"
	"os"
)

var Commands = []ucli.Command{
	{
		Name:    "blockchain",
		Aliases: []string{"bc"},
		Usage:   "Blockchain opertaions",
		Subcommands: []ucli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a blockchain and send genesis block reward to address",
				Flags: []ucli.Flag{
					ucli.StringFlag{
						Name: "address",
					},
				},
				Action: func(c *ucli.Context) error {
					if len(c.String("address")) < 1 {
						return ucli.ShowAppHelp(c)
					}

					createBlockChain(c.String("address"), os.Getenv("NODE_ID"))
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get all blockchain info",
				Action: func(c *ucli.Context) error {
					getBlockChain(os.Getenv("NODE_ID"))
					return nil
				},
			},
		},
	},
	{
		Name:    "transaction",
		Aliases: []string{"t"},
		Usage:   "Transaction opertaions",
		Subcommands: []ucli.Command{
			{
				Name:    "send",
				Aliases: []string{"s"},
				Usage:   "Send AMOUNT of coins from address A to address B",
				Flags: []ucli.Flag{
					ucli.StringFlag{
						Name: "from",
					},
					ucli.StringFlag{
						Name: "to",
					},
					ucli.IntFlag{
						Name: "amount",
					},
				},
				Action: func(c *ucli.Context) error {
					if len(c.String("from")) < 1 || len(c.String("to")) < 1 || c.Int("amount") <= 0 {
						return ucli.ShowAppHelp(c)
					}

					sendCoin(c.String("from"), c.String("to"), c.Int("amount"), os.Getenv("NODE_ID"), false)
					return nil
				},
			},
		},
	},
	{
		Name:    "wallet",
		Aliases: []string{"w"},
		Usage:   "Wallet opertaions",
		Subcommands: []ucli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create a new wallet",
				Action: func(c *ucli.Context) error {
					createWallet()
					return nil
				},
			},
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Get all wallets address",
				Action: func(c *ucli.Context) error {
					getWalletAddresses(os.Getenv("NODE_ID"))
					return nil
				},
			},
			{
				Name:    "balance",
				Aliases: []string{"b"},
				Usage:   "Get wallet balance info of address",
				Flags: []ucli.Flag{
					ucli.StringFlag{
						Name: "address",
					},
				},
				Action: func(c *ucli.Context) error {
					if len(c.String("address")) < 1 {
						return ucli.ShowAppHelp(c)
					}

					clog.Info("Address: %s, Balance: %d", c.String("address"), balance(c.String("address"), os.Getenv("NODE_ID")))
					return nil
				},
			},
		},
	},
	{
		Name:    "utxo",
		Aliases: []string{"u"},
		Usage:   "UTXO opertaions",
		Subcommands: []ucli.Command{
			{
				Name:    "reindex",
				Aliases: []string{"r"},
				Usage:   "Reindex the UTXO set",
				Action: func(c *ucli.Context) error {
					reindexUTXO(os.Getenv("NODE_ID"))
					return nil
				},
			},
		},
	},
}
