package main

import (
	"crypto/sha256"
	"time"
	"bytes"
	"encoding/binary"
)

//1.定义结构
type Block struct {
	PrevHash []byte //前区块哈希
	Hash []byte //当前区块哈希 实际的区块中是没有存储当前区块哈希的
	Data []byte //交易数据

//补充完成区块的结构
	Version uint64 //版本号
	MerkelRoot []byte //Merkel根 梅克尔跟 是一个哈希值 先不管 //TODO
	TimeStamp uint64 //时间戳
	Difficulty uint64 //难度值
	Nonce uint64 //随机数 也就是挖矿要找的数据
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
func NewBlock(data string,PrevHash []byte) *Block { //返回的是Block的指针
	block := Block{
		PrevHash:PrevHash,
		Hash:[]byte{},
		Data:[]byte(data),

		Version:00,
		MerkelRoot:[]byte{}, //TODO
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0, //随便填写的无效数据
		Nonce:0, //随便填写的无效数据
	}

	block.SetHash() //生成当前区块的哈希 用指针才能成功修改Hash的值

	return &block
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

	tmp := [][]byte {
		b.PrevHash,
		b.Data,
		Uint64ToByte(b.Version),
		b.MerkelRoot,
		Uint64ToByte(b.TimeStamp),
		Uint64ToByte(b.Difficulty),
		Uint64ToByte(b.Nonce),
	}

	var blockInfo []byte
	blockInfo = bytes.Join(tmp,[]byte{}) //将二维的切片数组连接起来 返回一个唯一的字符切片

	hash := sha256.Sum256(blockInfo) //将这个区块中所有的数据组成的信息生成哈希值
	b.Hash = hash[:] //把数组转成切片
}
