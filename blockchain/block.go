package blockchain

import (
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Timestamp     int64
	Data          string
	Nonce         int
	PrevBlockHash string
	Hash          string
}

/*func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)

	headers := bytes.Join([][]byte{[]byte(b.PrevBlockHash), data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = string(hash[:])
}*/

func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{time.Now().Unix(), data, 0, prevBlockHash, ""}
	//block.SetHash()
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
