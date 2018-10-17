package main

import (
	//"fmt"
)

//v2版本

//7.重构代码
func main()  {
	blockChain := NewBlockChain("THEO")
	cli := CLI{blockChain}
	cli.Run()

	//blockChain.AddBlock("吱吱兔给抓抓头转了50个比特币")
	//blockChain.AddBlock("吱吱兔又给抓抓头转了50个比特币")

/*
	for i,block := range blockChain.blocks {
		fmt.Printf("------------当前区块高度%d------------\n",i)
		fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
		fmt.Printf("当前哈希值:%x\n",block.Hash)
		fmt.Printf("区块数据:%s\n",block.Data)
		fmt.Printf("版本号：%b\n",block.Version)
		fmt.Printf("Merkel根：%s\n",block.MerkelRoot)
		fmt.Printf("时间戳：%b\n",block.TimeStamp)
		fmt.Printf("难度值：%b\n",block.Difficulty)
		fmt.Printf("随机数：%d\n",block.Nonce)
	}
*/
	//it := blockChain.NewBlockchainIterator()
	//for {
	//	block := it.Next()
	//
	//	fmt.Println("--------------------")
	//	fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
	//	fmt.Printf("当前哈希值:%x\n",block.Hash)
	//	fmt.Printf("区块数据:%s\n",block.Data)
	//	fmt.Printf("版本号：%b\n",block.Version)
	//	fmt.Printf("Merkel根：%s\n",block.MerkelRoot)
	//	fmt.Printf("时间戳：%b\n",block.TimeStamp)
	//	fmt.Printf("难度值：%b\n",block.Difficulty)
	//	fmt.Printf("随机数：%d\n",block.Nonce)
	//
	//	if len(block.PrevHash) == 0 {
	//		fmt.Println("区块链遍历结束")
	//		break
	//	}
	//}
}