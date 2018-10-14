package main

import "fmt"

//1.定义结构
type Block struct {
	PrevHash []byte //前区块哈希
	Hash []byte //当前区块哈希
	Data []byte //交易数据
}

//2.创建区块
func NewBlock(data string) Block {
	block := Block{
		PrevHash:[]byte{}, //TODO
		Hash:[]byte{}, //TODO
		Data:[]byte(data),
	}
	return block
}

//3.生成哈希
//4.引入区块链
//5.添加区块
//6.重构代码

func main()  {
	block := NewBlock("吱吱兔给抓抓头转了50个比特币")

	fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
	fmt.Printf("当前哈希值:%x\n",block.Hash)
	fmt.Printf("区块数据:%s\n",block.Data)

}