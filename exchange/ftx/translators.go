package ftx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/arbitrage/ticker"
)

func asTicker(
	source chan []byte,
	destination chan ticker.Ticker,
	ErrorC chan error,
) {
	var ftxTicker inMessage
	var ticker ticker.Ticker
	defer close(destination)

	for byteMessage := range source {
		if err := json.Unmarshal(byteMessage, &ftxTicker); err != nil {
			ErrorC <- fmt.Errorf("unmarshal err: %w", err)
			continue
		}

		ticker = ticker.Ticker{
			Exchange:   "FTX",
			Ask:        fmt.Sprintf("%.8f", ftxTicker.Data.Ask),
			Bid:        fmt.Sprintf("%.8f", ftxTicker.Data.Bid),
			AskSize:    fmt.Sprintf("%.8f", ftxTicker.Data.AskSize),
			BidSize:    fmt.Sprintf("%.8f", ftxTicker.Data.BidSize),
			Market:     ftxTicker.Market,
			Identifier: fmt.Sprintf("%v", time.Unix(int64(ftxTicker.Data.Time), 0)),
		}

		destination <- ticker
	}
}
