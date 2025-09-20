package main

import (
	"os"
	"os/signal"

	"github.com/domolitom/minotaur/pkg/binance"
	"github.com/domolitom/minotaur/pkg/detector"
)

func main() {
	det := detector.NewDetector(10, 300000) // 10 BTC, $300k

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go binance.RunOrderBookWS(det)
	go binance.RunTradeWS(det)

	<-interrupt
}
