package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type PestReport struct {
	Data         interface{}
	TimeStamp    time.Time
	PreviousHash string
	Hash         string
}

type ReportsRecord struct {
	Chain []PestReport
}

// CreateHash calculates and sets the hash for the block
func (block *PestReport) CreateHash() {
	newDataString, err := json.Marshal(block.Data)
	if err != nil {
		fmt.Println("error encoding data:", err)
		return
	}

	blockString := string(newDataString) + block.PreviousHash + block.TimeStamp.Format(time.RFC3339)

	blockHash := sha256.Sum256([]byte(blockString))

	block.Hash = hex.EncodeToString(blockHash[:])
}

// CreateGenesis initializes the genesis block and calculates its hash
func CreateGenesis() ReportsRecord {
	genesis := PestReport{
		Data:      "The first record",
		TimeStamp: time.Now(),
	}
	genesis.CreateHash()
	return ReportsRecord{
		Chain: []PestReport{genesis},
	}
}

// CreateBlock creates a new block and returns a new ReportsRecord with the new block appended
func (b ReportsRecord) CreateBlock() ReportsRecord {
	imgByte := []byte{0, 1} // Example data

	block := PestReport{
		Data:         imgByte,
		PreviousHash: b.Chain[len(b.Chain)-1].Hash,
		TimeStamp:    time.Now(),
	}

	block.CreateHash()

	newChain := append(b.Chain, block)
	return ReportsRecord{Chain: newChain}
}

// String provides a string representation of the blockchain for printing
func (b ReportsRecord) String() string {
	var result string
	for _, block := range b.Chain {
		dataStr, err := json.Marshal(block.Data)
		if err != nil {
			dataStr = []byte("error marshalling data")
		}
		result += fmt.Sprintf("TimeStamp: %s\nPreviousHash: %s\nHash: %s\nData: %s\n\n",
			block.TimeStamp.Format(time.RFC3339),
			block.PreviousHash,
			block.Hash,
			string(dataStr))
	}
	return result
}

func main() {
	blockchain := CreateGenesis()

	blockchain = blockchain.CreateBlock()

	fmt.Println(blockchain.String())
}
