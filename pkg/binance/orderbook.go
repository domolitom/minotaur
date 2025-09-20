package binance

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

type OrderBook struct {
	sync.RWMutex
	Bids map[string]decimal.Decimal // price => qty
	Asks map[string]decimal.Decimal
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make(map[string]decimal.Decimal),
		Asks: make(map[string]decimal.Decimal),
	}
}

func (ob *OrderBook) Update(side string, price string, qty string) {
	ob.Lock()
	defer ob.Unlock()
	amount, err := decimal.NewFromString(qty)
	if err != nil {
		return
	}
	if amount.IsZero() {
		if side == "bid" {
			delete(ob.Bids, price)
		} else {
			delete(ob.Asks, price)
		}
	} else {
		if side == "bid" {
			ob.Bids[price] = amount
		} else {
			ob.Asks[price] = amount
		}
	}
}

// Simple detection: print levels with amounts > threshold
func (ob *OrderBook) DetectJumps(thresh decimal.Decimal) {
	ob.RLock()
	defer ob.RUnlock()
	for price, qty := range ob.Bids {
		if qty.GreaterThan(thresh) {
			fmt.Printf("📈 BID JUMP: %s @ %s\n", qty, price)
		}
	}
	for price, qty := range ob.Asks {
		if qty.GreaterThan(thresh) {
			fmt.Printf("📈 ASK JUMP: %s @ %s\n", qty, price)
		}
	}
}
