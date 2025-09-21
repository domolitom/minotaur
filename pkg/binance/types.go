package binance

type DepthResult struct {
	Bids [][]string `json:"b"`
	Asks [][]string `json:"a"`
}

type DepthResponse struct {
	Stream string      `json:"stream"`
	Data   DepthResult `json:"data"`
}

type TradeEvent struct {
	Price        string `json:"p"`
	Qty          string `json:"q"`
	IsBuyerMaker bool   `json:"m"`
	TradeTime    int64  `json:"T"`
}
