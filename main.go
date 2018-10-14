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
func NewBlock(data string) *Block { //返回的是Block的指针
	block := Block{
		PrevHash:[]byte{}, //TODO
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
//5.添加区块
//6.重构代码

func main()  {
	block := NewBlock("吱吱兔给抓抓头转了50个比特币")

	fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
	fmt.Printf("当前哈希值:%x\n",block.Hash)
	fmt.Printf("区块数据:%s\n",block.Data)

}