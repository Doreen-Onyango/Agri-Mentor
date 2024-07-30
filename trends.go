package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Authority represents an authorized entity that can add blocks
type Authority struct {
	Name string
}

type Data struct {
	Location string
	Crop     string
	Price    float64
}

// AuthorizedAuthorities is a list of authorized nodes (authorities)
var AuthorizedAuthorities = map[string]Authority{
	"Authority1": {Name: "Authority1"},
	"Authority2": {Name: "Authority2"},
}

// PestReport represents a block in the blockchain
type PestReport struct {
	MarketInfo   Data
	TimeStamp    time.Time
	PreviousHash string
	Hash         string
	Authority    string // Track which authority created the block
}

// ReportsRecord represents the entire blockchain
type ReportsRecord struct {
	Chain []PestReport
}

// CreateHash calculates and sets the hash for the block
func (block *PestReport) CreateHash() {
	newDataString, err := json.Marshal(block.MarketInfo)
	if err != nil {
		fmt.Println("error encoding data:", err)
		return
	}

	blockString := string(newDataString) + block.PreviousHash + block.TimeStamp.Format(time.RFC3339) + block.Authority

	blockHash := sha256.Sum256([]byte(blockString))

	block.Hash = hex.EncodeToString(blockHash[:])
}

// CreateGenesis initializes the genesis block and calculates its hash
func CreateGenesis() ReportsRecord {
	genesis := PestReport{
		MarketInfo: Data{},
		TimeStamp:  time.Now(),
	}
	genesis.CreateHash()
	return ReportsRecord{
		Chain: []PestReport{genesis},
	}
}

// CreateBlock creates a new block and returns a new ReportsRecord with the new block appended
func (b ReportsRecord) CreateBlock(authorityName string, data Data) ReportsRecord {
	if _, authorized := AuthorizedAuthorities[authorityName]; !authorized {
		fmt.Println("Error: Unauthorized authority")
		return b
	}

	block := PestReport{
		MarketInfo:   data,
		PreviousHash: b.Chain[len(b.Chain)-1].Hash,
		TimeStamp:    time.Now(),
		Authority:    authorityName,
	}

	block.CreateHash()

	newChain := append(b.Chain, block)
	return ReportsRecord{Chain: newChain}
}

// String provides a string representation of the blockchain for printing
func (b ReportsRecord) String() string {
	var result string
	for _, block := range b.Chain {
		dataStr, err := json.Marshal(block.MarketInfo)
		if err != nil {
			dataStr = []byte("error marshalling data")
		}
		result += fmt.Sprintf("TimeStamp: %s\nPreviousHash: %s\nHash: %s\nData: %s\nAuthority: %s\n\n",
			block.TimeStamp.Format(time.RFC3339),
			block.PreviousHash,
			block.Hash,
			string(dataStr),
			block.Authority)
	}
	return result
}

var blockchain = CreateGenesis()

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain)
}

func addBlock(w http.ResponseWriter, r *http.Request) {
	var data Data
	authority := r.URL.Query().Get("authority")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	blockchain = blockchain.CreateBlock(authority, data)
	json.NewEncoder(w).Encode(blockchain)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/blockchain", getBlockchain)
	http.HandleFunc("/add", addBlock)
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
