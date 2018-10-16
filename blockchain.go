package main

import (
	"../bitcoin/bolt"
	"log"
	//"fmt"
)

//4.引入区块链
type BlockChain struct {
	//blocks []*Block //定义一个区块的数组 数组中是区块的指针

	db *bolt.DB //用数据库代替数组
    tail []byte //存储区块链中最后一个区块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//5.创建一个区块链
func NewBlockChain() *BlockChain  { //返回的是区块链的指针
	var lastHash []byte

	//1.打开数据库
	db,err := bolt.Open(blockChainDb,0600,nil) //打开数据库test.db 如果不存在就新建一个 0600是文件打开修改删除的权限模式
	if err != nil {
		log.Panic("数据库打开失败")
	}

	//2.操作数据库
	db.Update(func(tx *bolt.Tx) error {
		//找抽屉bucket 如果没有 就创建
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket,err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				panic("bucket blockBucket创建失败")
			}
			genesisBlock := NewBlock("创世块",[]byte{}) //创建创世块 并把它添加到区块链中 创世块的PrevHash是空

			//3.写数据
			bucket.Put(genesisBlock.Hash,genesisBlock.ToByte())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//fmt.Println(bucket.Get(genesisBlock.Hash))
			//fmt.Printf("%x",bucket.Get([]byte("lastHashKey")))
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})

	return &BlockChain{db:db,tail:lastHash}
}

//6.添加区块
func (bc *BlockChain) AddBlock(data string)  { //在区块链中添加区块
	//prevHash := bc.len(bc.blocks) - 1].Hash
	//newBlock := NewBlock(data,prevHash)
	//bc.blocks = append(bc.blocks,newBlock)
}
