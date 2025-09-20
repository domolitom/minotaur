package types

import "github.com/shopspring/decimal"

// Exchange-independent abstraction for an orderbook update
type OrderbookUpdate struct {
	Side  string // "bid" or "ask"
	Price decimal.Decimal
	Qty   decimal.Decimal
}

// Generic trade event
type TradeEvent struct {
	Price     decimal.Decimal
	Qty       decimal.Decimal
	Side      string // "buy" or "sell"
	Timestamp int64
	Exchange  string // e.g. "binance"
}
