package binance

type inTicker struct {
	Ask      string `json:"a"`
	Bid      string `json:"b"`
	AskSize  string `json:"A"`
	BidSize  string `json:"B"`
	UpdateId int    `json:"u"`
	Symbol   string `json:"s"`
}
