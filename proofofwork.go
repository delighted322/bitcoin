package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//在第一个版本中 区块的哈希值是无规则的

//1.定义一个工作量证明的结构
type ProofOfWork struct {
	block *Block
	target *big.Int //目标值 即目标哈希值 挖矿算出来的值要比这个值小才算挖矿成功 big.Int有着十分丰富的方法 如比较 赋值等
}

//2.提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork  {
	pow :=  ProofOfWork{
		block:block,
	}

	//我们指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	tmp := big.Int{}
	tmp.SetString(targetStr,16) //借助tmp这个big.Int类型的值将string转换成big.Int
									 //这个方法： 将targetStr转换成16进制 赋值给tmp

	pow.target = &tmp

	return &pow
}

//3.提供不断计算hash的函数
//(区块 + 随机数) -> 哈希值要小于目标哈希值 => 挖矿成功 生成区块
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1. 拼装数据（区块的数据，还有不断变化的随机数）
	b := pow.block
	var nonce uint64  //从默认值0开始试
	var hash [32]byte

	for {
		tmp := [][]byte {
			b.PrevHash,
			//b.Data, //只对区块头做哈希值 区块体通过MerkelRoot产生影响
			Uint64ToByte(b.Version),
			b.MerkelRoot,
			Uint64ToByte(b.TimeStamp),
			Uint64ToByte(b.Difficulty),
			Uint64ToByte(nonce),
		}

		var blockInfo []byte
		blockInfo = bytes.Join(tmp,[]byte{}) //将二维的切片数组连接起来 返回一个唯一的字符切片

		//2. 做哈希运算
		hash = sha256.Sum256(blockInfo) //将这个区块中所有的数据组成的信息生成哈希值

		//3. 与pow中的target进行比较
		tmpInt :=  big.Int{}
		tmpInt.SetBytes(hash[:]) //将hash[:]转换成big.Int{}并赋值给tmpInt

		if tmpInt.Cmp(pow.target) == -1 {
			//a. 找到了，退出返回
			fmt.Printf("挖矿成功: hash:%x，nonce：%d\n",hash[:],nonce)
			break
		} else {
			//b. 没找到，继续找，随机数加1
			nonce++
		}
	}



	return hash[:],nonce
}

//4.提供一个校验函数