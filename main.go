package main

import (
	"fmt"
	"crypto/sha256"
)

//1.定义结构
type Block struct {
	PrevHash []byte //前区块哈希
	Hash []byte //当前区块哈希
	Data []byte //交易数据
}

//2.创建区块
func NewBlock(data string,PrevHash []byte) *Block { //返回的是Block的指针
	block := Block{
		PrevHash:PrevHash,
		Hash:[]byte{},
		Data:[]byte(data),
	}

	block.SetHash() //生成当前区块的哈希 用指针才能成功修改Hash的值

	return &block
}

//3.生成哈希
func (b *Block) SetHash() {
	blockInfo := append(b.PrevHash,b.Data...) //拼装区块的数据
	hash := sha256.Sum256(blockInfo) //将这个区块中所有的数据组成的信息生成哈希值
	b.Hash = hash[:] //把数组转成切片
}

//4.引入区块链
type BlockChain struct {
	blocks []*Block //定义一个区块的数组 数组中是区块的指针
}

//5.创建一个区块链
func NewBlockChain() *BlockChain  { //返回的是区块链的指针
	genesisBlock := NewBlock("创世块",[]byte{}) //创建创世块 并把它添加到区块链中 创世块的PrevHash是空
	blockChain := BlockChain{blocks:[]*Block{genesisBlock}}
	return &blockChain
}

//6.添加区块
func (bc *BlockChain) AddBlock(data string)  { //在区块链中添加区块
	prevHash := bc.blocks[len(bc.blocks) - 1].Hash
	newBlock := NewBlock(data,prevHash)
	bc.blocks = append(bc.blocks,newBlock)
}

//7.重构代码

func main()  {
	blockChain := NewBlockChain()
	blockChain.AddBlock("吱吱兔给抓抓头转了50个比特币")
	blockChain.AddBlock("吱吱兔又给抓抓头转了50个比特币")


	for i,block := range blockChain.blocks {
		fmt.Printf("------------当前区块高度%d------------\n",i)
		fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
		fmt.Printf("当前哈希值:%x\n",block.Hash)
		fmt.Printf("区块数据:%s\n",block.Data)
	}
}