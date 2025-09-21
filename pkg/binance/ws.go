package binance

import (
	"encoding/json"
	"log"

	"github.com/domolitom/minotaur/pkg/detector"
	"github.com/domolitom/minotaur/pkg/types"
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
	c, _, err := websocket.DefaultDialer.Dial(tradesURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Trade error:", err)
			return
		}
		var trade TradeEvent
		if err := json.Unmarshal(msg, &trade); err != nil {
			continue
		}
		// Convert to generic detector.TradeEvent
		price, err1 := decimal.NewFromString(trade.Price)
		qty, err2 := decimal.NewFromString(trade.Qty)
		if err1 != nil || err2 != nil {
			continue
		}
		side := "sell"
		if !trade.IsBuyerMaker {
			side = "buy"
		}
		genericTrade := types.TradeEvent{
			Price:     price,
			Qty:       qty,
			Side:      side,
			Exchange:  "binance",
			Timestamp: trade.TradeTime,
		}
		det.DetectTrade(genericTrade)
	}
}
