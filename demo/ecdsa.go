package main

import (
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//演示如何使用ecdsa生成公钥私钥
//签名校验

func main()  {
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey,err := ecdsa.GenerateKey(curve,rand.Reader)
	if err != nil {
		log.Panic()
	}

	//生成公钥
	pubKey := privateKey.PublicKey

	data := "hello world!"
	hash := sha256.Sum256([]byte(data))

	//签名
	r,s,err := ecdsa.Sign(rand.Reader,privateKey,hash[:])
	if err != nil {
		log.Panic()
	}

	fmt.Printf("pubkey : %v\n", pubKey)
	fmt.Printf("r : %v, len : %d\n", r.Bytes(), len(r.Bytes()))
	fmt.Printf("s : %v, len : %d\n", s.Bytes(), len(s.Bytes()))

	//把r s进行序列化传输
	signature := append(r.Bytes(),s.Bytes()...)

	//定义两个辅助的big.Int
	r1 := big.Int{}
	s1 := big.Int{}

	//拆分signature 平均分 前半部分给r 后半部分给s
	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	//校验需要三个东西 数据签名 公钥
	res := ecdsa.Verify(&pubKey,hash[:],&r1,&s1)

	fmt.Printf("校验结果: %v\n",res)
}
