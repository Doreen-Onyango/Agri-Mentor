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
	timeStamp    time.Time
	previousHash string
	hash         string
}

type ReportsRecord struct {
	chain []PestReport
}

func (block *PestReport) CreateHash() {
	newDataString, err := json.Marshal(block.Data)
	if err != nil {
		fmt.Println("error encoding file")
		return
	}

	blockString := string(newDataString) + block.previousHash + block.timeStamp.Format(time.RFC3339)

	blockHash := sha256.Sum256([]byte(blockString))

	block.hash = hex.EncodeToString(blockHash[:])
}

func CreateGenesis() ReportsRecord {
	Genesis := PestReport{
		Data:      "The first record",
		timeStamp: time.Now(),
	}
	Genesis.CreateHash()
	return ReportsRecord{
		chain: []PestReport{Genesis},
	}
}

func (b ReportsRecord) createBlock() ReportsRecord {
	imgByte := []byte{0, 1}

	block := PestReport{
		Data:         imgByte,
		previousHash: b.chain[len(b.chain)-1].hash,
		timeStamp:    time.Now(),
	}

	block.CreateHash()
	b.chain = append(b.chain, block)
	fmt.Println(b.chain)
	return b
}

// String provides a string representation of the blockchain for printing
func (b ReportsRecord) Stringout() string {
	var result string
	for _, block := range b.chain {
		dataStr, _ := json.Marshal(block.Data)
		result += fmt.Sprintf("TimeStamp: %s\nPreviousHash: %s\nHash: %s\nData: %s\n\n",
			block.timeStamp.Format(time.RFC3339),
			block.previousHash,
			block.hash,
			string(dataStr))
	}
	return result
}

func main() {
	blockchain := CreateGenesis()

	blockchain.createBlock()

	// fmt.Println(blockchain)

	// for _, block := range blockchain.chain{
	// 	fmt.Println(block.Data)
	// 	fmt.Println(block.timeStamp)
	// 	fmt.Println(block.previousHash)
	// 	fmt.Println(block.hash)

	// }

	fmt.Println(blockchain.Stringout())
}

