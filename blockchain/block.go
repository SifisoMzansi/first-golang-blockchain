package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	nonce    int
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})

}

//	func (b *Block) DeriveHash() {
//		info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//		hash := sha256.Sum256(info)
//		b.Hash = hash[:]
//	}
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}
