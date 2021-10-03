package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

/*type Data struct {
	BottleId int
	OwnerId  int
}*/

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block's relevant information with the new blocks
	hash := sha256.Sum256(info)
	// This will hash the information and return a hash
	b.Hash = hash[:]
	// This will set the hash of the block to the hash of the information
	// [:] refers to golang slices and this is how you get the underlying array of the slice
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Refers to pointer's syntax and this is how you create a pointer to a block
	block.DeriveHash()

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func initBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := initBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		// This will loop through the blocks and print the data
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}
