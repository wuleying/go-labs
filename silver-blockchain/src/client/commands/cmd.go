package commands

import (
	"github.com/urfave/cli"
)

/*
 * e.g.
 * ./silver-blockchain b -address=ADDRESS   Get balance info of address
 * ./silver-blockchain bc -address=ADDRESS  Create a blockchain and send genesis block reward to address
 * ./silver-blockchain bc -print=1          Print all blockchain info
 */

var Commands = []cli.Command{
	{
		Name:    "balance",
		Aliases: []string{"b"},
		Usage:   "Balance operation",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "address",
				Usage: "Get balance info of address",
				Value: "",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.ShowAppHelp(c)
			}

			getBalance(c.String("address"))
			return nil
		},
	},
	{
		Name:    "blockchain",
		Aliases: []string{"bc"},
		Usage:   "Blockchain operation",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "address",
				Usage: "Create a blockchain and send genesis block reward to address",
				Value: "",
			},
			cli.StringFlag{
				Name:  "print",
				Usage: "Print all blockchain info",
				Value: "",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return cli.ShowAppHelp(c)
			}

			if c.Bool("print") {
				printBlockChain()
			} else if len(c.String("address")) > 0 {
				createBlockChain(c.String("address"))
			} else {
				return cli.ShowAppHelp(c)
			}

			return nil
		},
	},
}
