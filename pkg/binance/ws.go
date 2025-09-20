package binance

import (
	"encoding/json"
	"log"

	"github.com/domolitom/minotaur/pkg/detector"
	"github.com/gorilla/websocket"
)

const (
	pair         = "btcusdt"
	orderBookURL = "wss://stream.binance.com:9443/ws/" + pair + "@depth"
	tradesURL    = "wss://stream.binance.com:9443/ws/" + pair + "@trade"
)

func RunOrderBookWS(ob *OrderBook, det *detector.Detector) {
	c, _, err := websocket.DefaultDialer.Dial(orderBookURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Depth error:", err)
			return
		}
		var event DepthEvent
		if err := json.Unmarshal(msg, &event); err != nil {
			continue
		}
		for _, b := range event.Bids {
			ob.Update("bid", b[0], b[1])
		}
		for _, a := range event.Asks {
			ob.Update("ask", a[0], a[1])
		}
		det.DetectOrderbook(ob)
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
		det.DetectTrade(trade)
	}
}
