package blockchain

type BlockChain struct {
	Blocks []*Block
}

func InitBlockChain() *BlockChain {
	var chain BlockChain
	chain.Blocks = append(chain.Blocks, ZeroBlock())
	return &chain
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
