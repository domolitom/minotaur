package detector

import (
	"fmt"

	"github.com/domolitom/minotaur/pkg/binance"
	"github.com/shopspring/decimal"
)

type Detector struct {
	LargeQty      decimal.Decimal
	LargeTradeUSD decimal.Decimal
}

func NewDetector(largeQty, largeTradeUsd int64) *Detector {
	return &Detector{
		LargeQty:      decimal.NewFromInt(largeQty),
		LargeTradeUSD: decimal.NewFromInt(largeTradeUsd),
	}
}

func (d *Detector) DetectOrderbook(ob *binance.OrderBook) {
	ob.DetectJumps(d.LargeQty)
}

func (d *Detector) DetectTrade(trade binance.TradeEvent) {
	qty, err1 := decimal.NewFromString(trade.Qty)
	price, err2 := decimal.NewFromString(trade.Price)
	if err1 != nil || err2 != nil {
		return
	}
	usdValue := qty.Mul(price)
	if usdValue.GreaterThan(d.LargeTradeUSD) {
		direction := "SELL"
		if !trade.IsBuyerMaker {
			direction = "BUY"
		}
		fmt.Printf("🐋 LARGE TRADE: %s %v @ %v = $%v\n", direction, qty, price, usdValue)
	}
}
