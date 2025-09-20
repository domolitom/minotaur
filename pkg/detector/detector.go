package detector

import (
	"fmt"

	"github.com/domolitom/minotaur/pkg/types"
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

func (d *Detector) DetectOrderbook(update types.OrderbookUpdate) {
	if update.Qty.GreaterThan(d.LargeQty) {
		fmt.Printf("JUMP: %s %s @ %s\n", update.Side, update.Qty, update.Price)
	}
}
func (d *Detector) DetectTrade(event types.TradeEvent) {
	usdValue := event.Price.Mul(event.Qty)
	if usdValue.GreaterThan(d.LargeTradeUSD) {
		fmt.Printf("🐋 LARGE TRADE: %s %v @ %v on %s = $%v\n",
			event.Side, event.Qty, event.Price, event.Exchange, usdValue)
	}
}
