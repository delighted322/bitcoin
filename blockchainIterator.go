package main

import (
	"./bolt"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewBlockchainIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db:bc.db,
		currentHashPointer:bc.tail,
	}
}

