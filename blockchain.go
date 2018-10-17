package main

import (
	"../bitcoin/bolt"
	"log"
	"fmt"
	"bytes"
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
func NewBlockChain(address string) *BlockChain  { //返回的是区块链的指针
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
			genesisBlock := GenesisBlock(address) //创建创世块 并把它添加到区块链中 创世块的PrevHash是空
			fmt.Printf("genesisBlock :%s\n", genesisBlock)

			//3.写数据
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//这是为了读数据测试，马上删掉
			//blockBytes := bucket.Get(genesisBlock.Hash)
			//block := Deserialize(blockBytes)
			//fmt.Printf("block info : %s\n", block)
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})

	return &BlockChain{db:db,tail:lastHash}
}

//定义一个创世块
func GenesisBlock(address string) *Block  {
	coinbase := NewCoinbaseTX(address,"创世块")
	return NewBlock([]*TransAction{coinbase},[]byte{})
}

//6.添加区块
func (bc *BlockChain) AddBlock(txs []*TransAction)  { //在区块链中添加区块
	//prevHash := bc.len(bc.blocks) - 1].Hash
	//newBlock := NewBlock(data,prevHash)
	//bc.blocks = append(bc.blocks,newBlock)

	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		//找抽屉bucket 如果没有 就创建
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			panic("bucket 不应该为空 请检查")
		}

		block := NewBlock(txs,lastHash)

		//3.写数据
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("lastHashKey"),block.Hash)

		bc.tail = block.Hash  //一定要记得更新内存中的tail

		return nil
	})
}

func (bc *BlockChain) PrintChain()  {  //TODO 使用bolt.ForEach遍历区块链 不会
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key-> value 进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}
			block := Deserialize(v)
			//fmt.Printf("key=%x, value=%s\n", k, v)
			fmt.Printf("=============== 区块高度: %d ==============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig) //TODO 不懂
			return nil
		})
		return nil
	})
}

//找到指定地址的所有utxo
func (bc *BlockChain) FindUTOs(address string) []TXOutput  {
	var UTXO []TXOutput
	//

	return UTXO
}
