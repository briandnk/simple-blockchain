package main

import (
	"fmt"
	"github.com/briandnk/simple-blockchain/blockchain"
	"strconv"
)

type CommandLine struct {
}

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block after Root")
	chain.AddBlock("Second Block after Root")
	chain.AddBlock("Third Block after Root")

	for _, block := range chain.Blocks {
		fmt.Println("-------")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
