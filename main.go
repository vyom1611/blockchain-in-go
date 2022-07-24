package blockchain_in_go

import (
	"bytes"
	"crypto/sha256"
)

// Block : Defining a basic schema for a block
type Block struct {
	Hash     []byte
	Data     []byte
	prevHash []byte
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

func main() {

}
