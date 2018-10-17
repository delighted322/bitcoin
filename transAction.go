package main

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

//2.提供创建交易方法

//3.创建挖矿交易

//4.根据交易调整程序
