package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"crypto/rand"
)

type Wallet struct {
	//私钥
	Private *ecdsa.PrivateKey
	//约定 这里的PubKey不存储原始的公钥 而是存储X和Y拼接的字符串 在校验端重新拆分 (参考r s传递)
	PubKey []byte
}

func NewWallet() *Wallet  {
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic()
	}

	//生成公钥
	pubKeyOrig := privateKey.PublicKey

	//拼接X Y
	pubKey := append(pubKeyOrig.X.Bytes(),pubKeyOrig.Y.Bytes()...)

	return &Wallet{Private:privateKey,PubKey:pubKey}


}
