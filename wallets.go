package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"crypto/elliptic"
	"log"
	"os"
)

const walletFile  = "wallet.dat"

//定义一个Wallets结构 它保存所有的wallet以及它的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

//创建方法 返回当前所有钱包的实例
func NewWallets() *Wallets  {
	//wallet := NewWallet()
	//address := wallet.NewAddress()

	var wallets Wallets
	wallets.WalletsMap = make(map[string]*Wallet)
	//wallets.WalletsMap[address] = wallet

	wallets.loadFile()

	return &wallets
}

//读取文件 把所有的wallet读出来
func (ws *Wallets) loadFile()  {
	//在读取之前 要确认文件是否存在 如果不存在 直接退出
	_,err := os.Stat(walletFile)
	if os.IsNotExist(err) {
		return
	}

	//读取内容
	content,err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	//解码
	//panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256()) //记得先注册一下

	decoder := gob.NewDecoder(bytes.NewReader(content))

	var wsLocal Wallets

	err = decoder.Decode(&wsLocal)
	if err != nil {
		log.Panic(err)
	}

	ws.WalletsMap = wsLocal.WalletsMap //对于结构来说 里面有map的 要指定赋值 不要在最外层直接赋值 ws = &wsLocal 这样不行
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

	gob.Register(elliptic.P256())  //要先告诉gob编码是elliptic.P256

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil { // 一定要注意校验
		log.Panic(err)
	}

	ioutil.WriteFile(walletFile,buffer.Bytes(),0600)
}

func (ws *Wallets) ListAllAddresses() []string  {
	var addresses []string

	//遍历钱包 将所有的key取出来返回
	for address := range ws.WalletsMap {
		addresses = append(addresses,address)
	}

	return addresses
}
