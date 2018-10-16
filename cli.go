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
	addBlock --data DATA     "add data to blockChain"
	printChain               "print all blockChain data" 
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
	default:
		fmt.Println("无效的命令")
		fmt.Println(Usage)
	}
}