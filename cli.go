package main

import (
	"os"
	"fmt"
	"strconv"
)

//用来接收命令行参数并且控制区块链操作

type CLI struct {
	bc *BlockChain
}

const Usage  = `
	printChain               "正向打印区块链"
	printChainR              "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"
	send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
	newWallet   "创建一个新的钱包(私钥公钥对)"
`

func (cli *CLI) Run()  { //为什么不直接blockChain.Run*()呢 //TODO
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "send":
		fmt.Println("转账开始....")
		if len(args) != 7 {
			fmt.Println("参数个数错误 请检查")
			fmt.Printf(Usage)
			return
		}
		//send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
		from := args[2]
		to := args[3]
		amount,_ := strconv.ParseFloat(args[4],64) //将字符串转换成float64
		miner := args[5]
		data := args[6]
		cli.Send(from,to,amount,miner,data)
	case "printChain":
		fmt.Println("打印区块")
		cli.PrintBlockChain()
	case "getBalance":
		//fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	case "newWallet":
		fmt.Println("创建新的钱包...")
		cli.NewWallet()

	default:
		fmt.Println("无效的命令")
		fmt.Println(Usage)
	}
}