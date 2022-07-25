package main

import (
	"blockchain-in-go/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x \n", block.PrevHash)
		fmt.Printf("Data of current block: %x \n", block.Data)
		fmt.Printf("Hash of current block: %x \n", block.Hash)

		//Implementing proof of work service
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
