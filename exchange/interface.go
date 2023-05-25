package exchange

import (
	"github.com/arbitrage/ticker"
	"github.com/arbitrage/wsclient/connection"
)

type Partner interface {
	GetTicker(market string) (chan ticker.Ticker, error)
	StopTicker(market string) error
	TickerName(base string, quote string) (string, error)

	connection.Connector
}
