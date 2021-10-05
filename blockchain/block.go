package blockchain

import (
	"gorm.io/gorm"
)

type Data struct {
	BottleId string
	OwnerId  string
}

type Block struct {
	gorm.Model
	Hash     []byte
	BottleId []byte
	OwnerId  []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(bottleId string, ownerID string, prevHash []byte) *Block {
	block := &Block{Hash: []byte{}, BottleId: []byte(bottleId), OwnerId: []byte(ownerID), PrevHash: prevHash, Nonce: 0}
	// Refers to pointer's syntax and this is how you create a pointer to a block
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data *Data) *Block {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data.BottleId, data.OwnerId, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
	return new
}

func Genesis() *Block {
	return CreateBlock("Genesis bottle", "Genesis owner", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
