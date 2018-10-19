package main

import (
	"fmt"
	"time"
)

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
		timeForamt := time.Unix(int64(block.TimeStamp),0).Format("2006-01-02 15:04:05")
		fmt.Printf("时间戳：%s\n",timeForamt)
		fmt.Printf("Merkel根：%x\n",block.MerkelRoot)  //%x
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
	utxos := cli.bc.FindUTXOs(address)

	total := 0.0
	for _,utxo := range utxos {
		total += utxo.Value
	}

	fmt.Printf("\"%s\"的余额为：%f\n", address, total)
}

func (cli *CLI) Send(from, to string,amount float64,miner,data string)  {
	//fmt.Printf("from : %s\n", from)
	//fmt.Printf("to : %s\n", to)
	//fmt.Printf("amount : %f\n", amount)
	//fmt.Printf("miner : %s\n", miner)
	//fmt.Printf("data : %s\n", data)

	fmt.Printf("data:%s\n",data)

	//创建挖矿交易
	coinbase := NewCoinbaseTX(miner,data)
	//创建一个普通交易
	tx := NewTransaction(from,to,amount,cli.bc)
	if tx == nil {
		fmt.Printf("无效交易")
		return
	}
	//添加到区块
	cli.bc.AddBlock([]*TransAction{coinbase,tx})
	fmt.Printf("转账成功")
}

func (cli *CLI) NewWallet()  {
	//wallet := NewWallet()
	//address := wallet.NewAddress()
	//fmt.Printf("私钥：%v\n", wallet.Private)
	//fmt.Printf("公钥：%v\n", wallet.PubKey)
	//fmt.Printf("地址：%s\n", address)

	ws := NewWallets()
	for address := range ws.WalletsMap {
		fmt.Printf("地址：%s\n",address)
	}
}