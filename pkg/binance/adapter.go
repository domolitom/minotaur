package binance

import (
	"github.com/domolitom/minotaur/pkg/types"

	"github.com/shopspring/decimal"
)

// Converts Binance-specific trade event to generic TradeEvent
func ToGenericTrade(e TradeEvent) types.TradeEvent {
	price, _ := decimal.NewFromString(e.Price)
	qty, _ := decimal.NewFromString(e.Qty)
	side := "sell"
	if !e.IsBuyerMaker {
		side = "buy"
	}
	return types.TradeEvent{
		Price:     price,
		Qty:       qty,
		Side:      side,
		Timestamp: e.TradeTime,
		Exchange:  "binance",
	}
}
