package binance

import (
	"github.com/arbitrage/apis/binance/rest"
)

func New() Binance {
	// TODO: Add apiKey and apiSecret here
	return Binance{
		Rest: rest.Rest{},
	}
}
