package blockchain

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

func ZeroBlock() *Block {
	return CreateBlock("Zero Block", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:         []byte(data),
		PreviousHash: prevHash,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
