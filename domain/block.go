package domain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Index         int32
	Difficulty    int32
	Nonce         int32
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {

	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func newBlock(data string, PrevBlockHash []byte) *Block {

	block := &Block{
		Index:         1,
		Difficulty:    2,
		Nonce:         1,
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{},
	}

	block.SetHash()

	return block
}
