package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dimensi0n/chainpagne/blockchain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	chain := blockchain.InitBlockChain()

	db, err := gorm.Open(sqlite.Open("tmp/blocks.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&blockchain.Block{})
	var genesis blockchain.Block
	err = db.First(&genesis, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No blocks in database")
		db.Create(chain.Blocks[0])
		db.Create(chain.AddBlock(&blockchain.Data{BottleId: "01", OwnerId: "01"}))
		db.Create(chain.AddBlock(&blockchain.Data{BottleId: "02", OwnerId: "02"}))
		db.Create(chain.AddBlock(&blockchain.Data{BottleId: "01", OwnerId: "02"}))
	} else {
		fmt.Println("Blocks in database")
		var blocks []blockchain.Block
		db.Find(&blocks)

		for _, databaseBlock := range blocks {
			chain.AddBlock(&blockchain.Data{BottleId: string(databaseBlock.BottleId), OwnerId: string(databaseBlock.OwnerId)})
		}
	}

	for _, block := range chain.Blocks {
		// This will loop through the blocks and print the data
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Bottle ID: %s\n", block.BottleId)
		fmt.Printf("Owner ID: %s\n", block.OwnerId)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
