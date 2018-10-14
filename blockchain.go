package main

//4.引入区块链
type BlockChain struct {
	blocks []*Block //定义一个区块的数组 数组中是区块的指针
}

//5.创建一个区块链
func NewBlockChain() *BlockChain  { //返回的是区块链的指针
	genesisBlock := NewBlock("创世块",[]byte{}) //创建创世块 并把它添加到区块链中 创世块的PrevHash是空
	blockChain := BlockChain{blocks:[]*Block{genesisBlock}}
	return &blockChain
}

//6.添加区块
func (bc *BlockChain) AddBlock(data string)  { //在区块链中添加区块
	prevHash := bc.blocks[len(bc.blocks) - 1].Hash
	newBlock := NewBlock(data,prevHash)
	bc.blocks = append(bc.blocks,newBlock)
}
