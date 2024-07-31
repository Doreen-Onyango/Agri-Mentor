package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"agri-mentor/agri"
)

type Service struct {
	Icon        string
	Title       string
	Description string
}

type PageData struct {
	Services []Service
}

type MarketInfo struct {
	Location string  `json:"Location"`
	Crop     string  `json:"Crop"`
	Price    float64 `json:"Price"`
}

type Block struct {
	MarketInfo   MarketInfo `json:"MarketInfo"`
	TimeStamp    time.Time  `json:"TimeStamp"`
	PreviousHash string     `json:"PreviousHash"`
	Hash         string     `json:"Hash"`
	Authority    string     `json:"Authority"`
}

type Blockchain struct {
	Chain []Block `json:"Chain"`
}

var blockchain = createGenesis()

var AuthorizedAuthorities = map[string]bool{
	"Authority1": true,
	"Authority2": true,
	// Add more authorized authorities as needed
}

func createHash(block Block) string {
	record := block.MarketInfo.Location + block.MarketInfo.Crop + fmt.Sprintf("%f", block.MarketInfo.Price) + block.PreviousHash + block.TimeStamp.String() + block.Authority
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createGenesis() Blockchain {
	genesisBlock := Block{
		MarketInfo:   MarketInfo{},
		TimeStamp:    time.Now(),
		PreviousHash: "",
		Authority:    "Genesis",
	}
	genesisBlock.Hash = createHash(genesisBlock)
	return Blockchain{Chain: []Block{genesisBlock}}
}

func (bc *Blockchain) addBlock(marketInfo MarketInfo, authority string) {
	newBlock := Block{
		MarketInfo:   marketInfo,
		TimeStamp:    time.Now(),
		PreviousHash: bc.Chain[len(bc.Chain)-1].Hash,
		Authority:    authority,
	}
	newBlock.Hash = createHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
}

func main() {
	// Serve static files from the /static directory
	//http.Handle("/", http.FileServer(http.Dir("./templates")))

	// http.HandleFunc("/", server.RenderTemplate) // Serve the HTML file

	http.HandleFunc("/query", agri.HandleSendMessage) // Handle the input submission

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/blockchain", getBlockchain)
	http.HandleFunc("/add", addBlockHandler)
	http.HandleFunc("/market-trends", marketTrendsHandler)

	log.Println("Server starting on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Services: []Service{
			{Icon: "fa-seedling", Title: "Crop Management", Description: "Get personalized advice on rice cultivation techniques."},
			{Icon: "fa-bug", Title: "Pest Control", Description: "Identify and manage pests affecting your rice crops."},
			{Icon: "fa-chart-line", Title: "Market Trends", Description: "Stay updated on rice market prices and trends."},
			{Icon: "fa-cloud-sun-rain", Title: "Weather Forecasts", Description: "Get accurate weather predictions for your farm location."},
		},
	}

	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain)
}

func addBlockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var marketInfo MarketInfo
	err := json.NewDecoder(r.Body).Decode(&marketInfo)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	authority := r.URL.Query().Get("authority")
	if authority == "" {
		http.Error(w, "Authority is required", http.StatusBadRequest)
		return
	}

	if !AuthorizedAuthorities[authority] {
		http.Error(w, "Unauthorized authority", http.StatusUnauthorized)
		return
	}

	blockchain.addBlock(marketInfo, authority)
	w.WriteHeader(http.StatusCreated)
}

func marketTrendsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/market-trends.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Authorities []string
	}{
		Authorities: getAuthorizedAuthorities(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getAuthorizedAuthorities() []string {
	authorities := make([]string, 0, len(AuthorizedAuthorities))
	for auth := range AuthorizedAuthorities {
		authorities = append(authorities, auth)
	}
	return authorities
}
