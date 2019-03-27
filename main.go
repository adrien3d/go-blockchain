package main

import (
	"fmt"
	"github.com/adrien3d/monarch/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Adrien")
	bc.AddBlock("Send 2 more BTC to Adrien")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\t Nonce: %d\n", strconv.FormatBool(pow.Validate()), block.Nonce)
	}
}
