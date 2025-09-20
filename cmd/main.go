package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// BinancePriceResponse represents the relevant field from the API response
type BinancePriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type DepthResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

func main() {
	http.HandleFunc("/price", priceHandler)
	http.HandleFunc("/orderbook", orderBookHandler)
	fmt.Println("Starting server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// GetPrice fetches the current price for a given trading pair from Binance
func GetPrice(symbol string) (BinancePriceResponse, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return BinancePriceResponse{}, err
	}
	defer resp.Body.Close()

	var priceResp BinancePriceResponse
	err = json.NewDecoder(resp.Body).Decode(&priceResp)
	return priceResp, err
}

func GetOrderBook(symbol string, limit int) (DepthResponse, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=%d", symbol, limit)
	resp, err := http.Get(url)
	if err != nil {
		return DepthResponse{}, err
	}
	defer resp.Body.Close()

	var depth DepthResponse
	err = json.NewDecoder(resp.Body).Decode(&depth)
	return depth, err
}

// HTTP handler
func priceHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		symbol = "BTCUSDT"
	}
	priceResp, err := GetPrice(symbol)
	if err != nil {
		http.Error(w, "Failed to fetch price from Binance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(priceResp)
}

func orderBookHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		symbol = "BTCUSDT"
	}
	limit := 5 // Default limit
	depthResp, err := GetOrderBook(symbol, limit)
	if err != nil {
		http.Error(w, "Failed to fetch order book from Binance", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(depthResp)
}
