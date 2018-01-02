package main

import (
	"fmt"
	b "go-labs/silver-blockchain/src/block"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	bc := b.NewBlockchain()

	bc.AddBlock("First block")
	bc.AddBlock("Second block")

	for _, block := range bc.Blocks {
		fmt.Printf("Id: #%d\n", block.Id)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := b.NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	fmt.Printf("Run elapsed: %s\n", time.Since(startTime))
}
