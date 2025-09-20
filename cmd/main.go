package main

import (
	"os"
	"os/signal"

	"github.com/domolitom/minotaur/pkg/binance"
	"github.com/domolitom/minotaur/pkg/detector"
)

func main() {
	ob := binance.NewOrderBook()
	det := detector.NewDetector(10, 300000) // 10 BTC, $300k

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go binance.RunOrderBookWS(ob, det)
	go binance.RunTradeWS(det)

	<-interrupt
}
