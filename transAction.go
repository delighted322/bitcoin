package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

const reward  = 50 //每次挖矿成功得到的奖励

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
	//Sig string

	//真正的数字签名 由r s拼成的[]byte
	Signature []byte

	//约定 这里的PubKey不存储原始的公钥 而是存储X和Y拼接的字符串 在校验端重新拆分 (参考r s传递)
	PubKey []byte
}

type TXOutput struct {
	//转账金额
	Value float64 //要大写
	//锁定脚本(对方公钥的哈希 这个哈希可以通过地址反推出来 所以转账时知道地址即可) 这里用地址模拟
	//PubKeyHash string

	//收款方的公钥的哈希 注意 是哈希而不是公钥 也不是地址
	PubKeyHash []byte
}

//由于现在存储的字段是地址的公钥哈希 所以无法直接创建TXOutput
//为了能够得到公钥哈希 我们需要处理一下 写一个Lock函数
func (output *TXOutput) Lock(address string)  {
	//1.解码
	//2.截取出公钥哈希 去除version(1字节)去除校验码(4字节)
	addressByte := base58.Decode(address) //25字节
	len := len(addressByte)

	pubKeyHash := addressByte[1:(len-4)]

	//真正的锁定动作
	output.PubKeyHash = pubKeyHash
}

func NewTxOutput(value float64,address string) *TXOutput  {
	output := TXOutput{
		Value:value,
	}
	output.Lock(address)
	return &output
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

//实现一个函数 判断当前的交易是否为挖矿交易
func (tx *TransAction) IsCoinbase() bool {
	//交易input只有一个
	if len(tx.TXInputs) == 1 {
		input := tx.TXInputs[0]
		//交易id为空
		//交易的inde为-1
		if bytes.Equal(input.TXid,[]byte{}) && input.Index == -1 {
			return true
		}
	}

	return false
}

//创建普通的转账交易

//创建outputs

//如果有零钱 要找零
func NewTransaction(from, to string, amount float64, bc *BlockChain) *TransAction  {
	//找到最合理的UTXO集合 map[string][]uint64
	utxos,resValue := bc.FindNeedUTXOs(from,amount)

	if resValue < amount {
		fmt.Println("余额不足 交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	//创建交替输入 将这些utxo逐一转成inputs
	for id, indexArray := range utxos {
		for _,i := range indexArray {
			input := TXInput{[]byte(id),int64(i),from}
			inputs = append(inputs,input)
		}
	}

	//创建交易输出
	output := TXOutput{amount,to}
	outputs = append(outputs,output)

	//找零
	if resValue > amount {
		outputs = append(outputs,TXOutput{resValue-amount,from})
	}

	tx := TransAction{[]byte{},inputs,outputs}
	tx.SetHash()
	return  &tx
}

//4.根据交易调整程序
