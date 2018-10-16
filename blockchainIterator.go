package main

import (
	"./bolt"
	"log"
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

//1.返回当前区块
//2.指针前移
func (it *BlockChainIterator) Next() *Block  {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时bucket不应该为空，请检查!")
		}
		tmpBlock := bucket.Get([]byte(it.currentHashPointer))
		block = Deserialize(tmpBlock)
		it.currentHashPointer = block.PrevHash

		return nil
	})

	return &block
}

