package binance

import (
	"fmt"
	"log"

	"github.com/domolitom/minotaur/pkg/detector"
	"github.com/gorilla/websocket"
	"github.com/shopspring/decimal"
)

const (
	pair         = "btcusdt"
	orderBookURL = "wss://stream.binance.com:9443/ws/" + pair + "@depth"
	tradesURL    = "wss://stream.binance.com:9443/ws/" + pair + "@trade"
)

func RunOrderBookWS(det *detector.Detector) {
	conn, _, err := websocket.DefaultDialer.Dial(orderBookURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	var (
		ob  = NewOrderBook()
		res DepthResponse
	)

	for {
		if err := conn.ReadJSON(&res); err != nil {
			log.Fatal("Orderbook read:", err)
		}
		ob.handleDepthResponse(res.Data)
	}
}

func RunTradeWS(det *detector.Detector) {
	conn, _, err := websocket.DefaultDialer.Dial(tradesURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	var res TradeResponse

	for {
		if err := conn.ReadJSON(&res); err != nil {
			log.Fatal("Trade read:", err)
		}
		price, err := decimal.NewFromString(res.Data.Price)
		if err != nil {
			log.Println("Trade price parse:", err)
			continue
		}
		qty, err := decimal.NewFromString(res.Data.Qty)
		if err != nil {
			log.Println("Trade qty parse:", err)
			continue
		}
		fmt.Println("Trade:", price, qty, res.Data.IsBuyerMaker, res.Data.TradeTime)
	}
}
