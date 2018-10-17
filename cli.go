package main

import (
	"os"
	"fmt"
)

//用来接收命令行参数并且控制区块链操作

type CLI struct {
	bc *BlockChain
}

const Usage  = `
	addBlock --data DATA     "添加区块"
	printChain               "正向打印区块链"
	printChainR              "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"	
`

func (cli *CLI) Run()  { //为什么不直接blockChain.Run*()呢 //TODO
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不当，请检查")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Println("打印区块")
		cli.PrintBlockChain()
	case "getBalance":
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	default:
		fmt.Println("无效的命令")
		fmt.Println(Usage)
	}
}