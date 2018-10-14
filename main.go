package main

import (
	"fmt"
)

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