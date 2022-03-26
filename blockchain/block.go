package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func ZeroBlock() *Block {
	return CreateBlock("Zero Block", []byte{})
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:         []byte(data),
		PreviousHash: prevHash,
	}
	block.DeriveHash()

	return block
}
