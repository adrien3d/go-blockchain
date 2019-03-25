package main

import (
	"fmt"
	"strconv"
)

type Blockchain struct {
	blocks []*Block
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Adrien")
	bc.AddBlock("Send 2 more BTC to Adrien")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\t Nonce: %d\n", strconv.FormatBool(pow.Validate()), block.Nonce)
		fmt.Println()
	}
}
