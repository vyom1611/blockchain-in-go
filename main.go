package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block : Defining a basic schema for a block
type Block struct {
	Hash     []byte
	Data     []byte
	prevHash []byte
}

//BlockChain : Created schema for a blockchain which is a simple array for block types
type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.prevHash}, []byte{})
	//Generating hash using SHA256 algorithm
	hash := sha256.Sum256(info)
	//Setting the generated hash to the block's hash property
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	//Creating hash code
	block.DeriveHash()
	return block
}

//Adding block to array blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)

	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//Initialising blockchain with creating the first block with genesis
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x \n", block.prevHash)
		fmt.Printf("Data of current block: %x \n", block.Data)
		fmt.Printf("Hash of current block: %x \n", block.Hash)
	}
}
