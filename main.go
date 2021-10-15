package main

import (
	"fmt"

	"github.com/ashercarson/gocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Blcok")
	chain.AddBlock("Third Blcok")
	chain.AddBlock("Fourth Blcok")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	}
}