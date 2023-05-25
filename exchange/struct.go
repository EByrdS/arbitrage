package exchange

import (
	"github.com/arbitrage/ticker"
	"github.com/arbitrage/wsclient"
)

type Exchange struct {
	*wsclient.WSClient

	Name    string
	Tickers map[string]chan ticker.Ticker
}
