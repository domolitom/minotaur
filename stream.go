package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
