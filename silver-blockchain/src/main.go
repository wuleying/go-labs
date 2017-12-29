package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("First block.")
	bc.AddBlock("Second block.")

	for _, block := range bc.blocks {
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
