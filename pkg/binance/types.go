package binance

type DepthEvent struct {
	Bids [][]string `json:"b"`
	Asks [][]string `json:"a"`
}

type TradeEvent struct {
	Price        string `json:"p"`
	Qty          string `json:"q"`
	IsBuyerMaker bool   `json:"m"`
	TradeTime    int64  `json:"T"`
}
