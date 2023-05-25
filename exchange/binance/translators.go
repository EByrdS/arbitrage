package binance

import (
	"encoding/json"
	"fmt"

	"github.com/arbitrage/ticker"
)

func asTicker(
	source chan []byte,
	destination chan ticker.Ticker,
	ErrorC chan error,
) {
	var binanceTicker inTicker
	var ticker ticker.Ticker
	defer close(destination)

	for byteMessage := range source {
		if err := json.Unmarshal(byteMessage, &binanceTicker); err != nil {
			ErrorC <- fmt.Errorf("unmarshal err: %w", err)
			continue
		}

		ticker = ticker.Ticker{
			Exchange:   "Binance",
			Ask:        binanceTicker.Ask,
			Bid:        binanceTicker.Bid,
			AskSize:    binanceTicker.AskSize,
			BidSize:    binanceTicker.BidSize,
			Market:     binanceTicker.Symbol,
			Identifier: fmt.Sprintf("%d", binanceTicker.UpdateId),
		}

		destination <- ticker
	}
}
