package main

import (
	"fmt"

	"github.com/orrrlov/blockchain-go/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
