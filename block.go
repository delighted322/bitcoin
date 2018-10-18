package main

import (
	//"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//1.定义结构
type Block struct {
	PrevHash []byte //前区块哈希
	Hash []byte //当前区块哈希 实际的区块中是没有存储当前区块哈希的
	//Data []byte //交易数据
	Transactions []*TransAction //真实的交易数据 一个区块里可能有多个交易

//补充完成区块的结构
	Version uint64 //版本号
	MerkelRoot []byte //Merkel根 梅克尔跟 是一个哈希值 先不管 //TODO
	TimeStamp uint64 //时间戳
	Difficulty uint64 //难度值
	Nonce uint64 //随机数 也就是挖矿要找的数据
}

//序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错")
	}
	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block  {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错!")
	}
	return block
}

//将uint64类型转换成byte类型
func Uint64ToByte(num uint64) []byte { //TODO 不懂
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

//2.创建区块
func NewBlock(txs []*TransAction,PrevHash []byte) *Block { //返回的是Block的指针
	block := Block{
		PrevHash:PrevHash,
		Hash:[]byte{},
		//Data:[]byte(data),
		Transactions:txs,

		Version:00,
		//MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0, //随便填写的无效数据
		Nonce:0, //随便填写的无效数据
	}

	block.MerkelRoot = block.MakeMerkelRoot()

	//block.SetHash() //生成当前区块的哈希 用指针才能成功修改Hash的值
	pow := NewProofOfWork(&block)
	hash,nonce := pow.Run() //在Run()中不断进行哈希运算 直到找到一个随机数 使得生成的哈希值小于目标哈希值 则返回当前哈希 以及这个随机数

	block.Hash = hash
	block.Nonce = nonce

	return &block
}

//模拟梅克尔根 只是对交易的数据进行简单的拼接 而不做二叉树处理
func (b *Block) MakeMerkelRoot() []byte{
	var info []byte
	for _,tx := range b.Transactions {
		//将交易的哈希值(交易ID)拼接起来 再整体做哈希处理
		info = append(info,tx.TXID...)
	}
	hash := sha256.Sum256(info)
	return hash[:]
}

//3.生成哈希
func (b *Block) SetHash() {
	/*
	blockInfo := append(b.PrevHash,b.Data...) //拼装区块的数据
	blockInfo = append(blockInfo,Uint64ToByte(b.Version)...)
	blockInfo = append(blockInfo,b.MerkelRoot...)
	blockInfo = append(blockInfo,Uint64ToByte(b.TimeStamp)...)
	blockInfo = append(blockInfo,Uint64ToByte(b.Difficulty)...)
	blockInfo = append(blockInfo,Uint64ToByte(b.Nonce)...)
	*/
	//
	//tmp := [][]byte {
	//	b.PrevHash,
	//	b.Data,
	//	Uint64ToByte(b.Version),
	//	b.MerkelRoot,
	//	Uint64ToByte(b.TimeStamp),
	//	Uint64ToByte(b.Difficulty),
	//	Uint64ToByte(b.Nonce),
	//}
	//
	//var blockInfo []byte
	//blockInfo = bytes.Join(tmp,[]byte{}) //将二维的切片数组连接起来 返回一个唯一的字符切片
	//
	//hash := sha256.Sum256(blockInfo) //将这个区块中所有的数据组成的信息生成哈希值
	//b.Hash = hash[:] //把数组转成切片
}
