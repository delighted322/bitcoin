package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

const reward  = 12.5 //每次挖矿成功得到的奖励

//1.定义交易结构
type TransAction struct {
	TXID []byte //交易ID 一般是交易结构的哈希值
	TXInputs []TXInput //交易输入数组
	TXOutputs []TXOutput //交易输出的数组
}

type TXInput struct {
	//引用utxo所在的交易ID
	TXid []byte
	//所消费utxo在output中的索引
	Index int64
	//解锁脚本(签名 公钥) 我们用地址来模拟
	Sig string
}

type TXOutput struct {
	//转账金额
	value float64
	//锁定脚本(对方公钥的哈希 这个哈希可以通过地址反推出来 所以转账时知道地址即可) 这里用地址模拟
	PubKeyHash string
}

//设置交易ID
func (tx *TransAction) SetHash() {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err,"编码出错")
	}

	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}


//2.提供创建交易的方法(挖矿交易)
func NewCoinbaseTX(address string,data string) *TransAction{ //address模拟锁定脚本 data是由矿工填写的sig字段
	//挖矿交易的特点
	//1.只有一个输入交易
	//2.无需引用交易id
	//3.无需引用index
	//矿工由于挖矿时无需指定签名 所以这个sig字段可以由矿工自由填写数据 一般是填写矿池的名字
	input := TXInput{[]byte{},-1,data}
	output := TXOutput{reward,address}

	//对于挖矿交易来说 只有一个input和一个output
	tx := TransAction{TXID:[]byte{},TXInputs:[]TXInput{input},TXOutputs:[]TXOutput{output}}
	tx.SetHash() //生成交易的ID

	return &tx
}

//3.创建挖矿交易

//4.根据交易调整程序
