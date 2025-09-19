package main

import (
	"fmt"
	"log"
	"net/http"
)

// BinancePriceResponse represents the relevant field from the API response
type BinancePriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	http.HandleFunc("/price", priceHandler)
	fmt.Println("Starting server :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
