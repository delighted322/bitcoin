package main

import "math/big"

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
func (pow *ProofOfWork) Run() (hash []byte, nonce uint64) {
	//TODO
	return []byte("helloWorld"),0
}

//4.提供一个校验函数