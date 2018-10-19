package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

//定义一个Wallets结构 它保存所有的wallet以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

func NewWallets() *Wallets  {
	//wallet := NewWallet()
	//address := wallet.NewAddress()

	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	//wallets.WalletsMap[address] = wallet

	return &wallets
}

func (ws *Wallets) CreateWallet() string  {
	wallet := NewWallet()
	address := wallet.NewAddress()

	ws.WalletsMap[address] = wallet

	ws.saveToFile()
	return address
}

//保存方法 把新建的wallet添加进去
func (ws *Wallets) saveToFile()  {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(ws)

	ioutil.WriteFile("wallet.dat",buffer.Bytes(),0600)
}
