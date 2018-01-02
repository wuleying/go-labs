package main

import (
	b "go-labs/silver-blockchain/src/block"
)

func main() {
	bc := b.NewBlockchain()
	defer bc.Db.Close()

	cli := CLI{bc}
	cli.Run()
}
