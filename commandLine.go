package main

import "fmt"

func (cli *CLI) AddBlock(data string)  {
	//cli.bc.AddBlock(data) //TODO
	fmt.Println("添加区块成功")
}

func (cli *CLI) PrintBlockChain()  { //TODO
	blockChain := cli.bc

	it := blockChain.NewBlockchainIterator()
	for {
		block := it.Next()

		fmt.Println("--------------------")
		fmt.Printf("前区块哈希值:%x\n",block.PrevHash)
		fmt.Printf("当前哈希值:%x\n",block.Hash)
		fmt.Printf("区块数据:%s\n",block.Transactions[0].TXInputs[0].Sig)
		fmt.Printf("版本号：%b\n",block.Version)
		fmt.Printf("Merkel根：%s\n",block.MerkelRoot)
		fmt.Printf("时间戳：%b\n",block.TimeStamp)
		fmt.Printf("难度值：%b\n",block.Difficulty)
		fmt.Printf("随机数：%d\n",block.Nonce)

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}
}

//获取指定地址的余额
func (cli *CLI) GetBalance(address string)  {
	utxos := cli.bc.FindUTOs(address)

	total := 0.0
	for _,utxo := range utxos {
		total += utxo.value
	}

	fmt.Printf("\"%s\"的余额为：%f\n", address, total)
}
