package main

import "fmt"
import (
	b "go-labs/silver-blockchain/src/block"
	"strconv"
)

func main() {
	bc := b.NewBlockchain()

	bc.AddBlock("First block.")
	bc.AddBlock("Second block.")

	for _, block := range bc.Blocks {
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := b.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}