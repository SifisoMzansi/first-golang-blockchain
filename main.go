package main

import (
	"fmt"
	"strconv"

	"github.com/SifisoMzansi/first-golang-blockchain/blockchain"
)

func main() {

	chain := blockchain.InitBlockChain()
	chain.AddBlock("Block 1 after genesis ")
	chain.AddBlock("Block 2 after genesis ")
	chain.AddBlock("Block 3 after genesis ")

	for _, block := range chain.Blocks {
		fmt.Printf("data: %s\n ", block.Data)
		fmt.Printf("hash: %x\n ", block.Hash)
		fmt.Printf("Previous hash: %x\n ", block.PrevHash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Proof of work is: %s\n ", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	// for i := 0; i < len(chain.Blocks); i++ {
	// 	fmt.Println("Count ", i)
	// 	fmt.Printf("data: %s\n ", chain.Blocks[i].Data)
	// 	fmt.Printf("hash: %x\n ", chain.Blocks[i].Hash)
	// 	fmt.Printf("Previous hash: %x\n ", chain.Blocks[i].PrevHash)
	// }

	// for _, block := range chain.Blocks {
	// 	fmt.Printf("data: %s\n ", block.Data)
	// 	fmt.Printf("hash: %x\n ", block.Hash)
	// 	fmt.Printf("Previous hash: %x\n ", block.PrevHash)
	// }
}
