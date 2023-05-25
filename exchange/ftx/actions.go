package ftx

import (
	"encoding/json"
	"fmt"

	"github.com/arbitrage/ticker"
)

func tickerChanName(market string) string {
	return "ticker-" + market
}

// same signature as lane.Lane.Belongs
// must map []byte to expected structure in market
// creating a closure
func tickerBelongs(market string) func([]byte) bool {
	return func(byteMessage []byte) bool {
		var received inMessage

		if err := json.Unmarshal(byteMessage, &received); err != nil {
			return false
		}

		return received.Type == "update" && received.Channel == "ticker" && received.Market == market
	}
}

func (f FTX) GetTicker(market string) (chan ticker.Ticker, error) {
	var tickerChan chan ticker.Ticker
	msg := outMessage{
		Op:      "subscribe",
		Channel: "ticker",
		Market:  market,
	}

	channelName := tickerChanName(market)

	byteLane, err := f.Subscribe(msg, channelName, tickerBelongs(market))
	if err != nil {
		return tickerChan, fmt.Errorf("subscribe error: %w", err)
	}

	destination := make(chan ticker.Ticker)
	f.Tickers[channelName] = destination
	go asTicker(byteLane.C, destination, f.ErrorC())

	return destination, nil
}

func (f FTX) StopTicker(market string) error {
	msg := outMessage{
		Op:      "unsubscribe",
		Channel: "ticker",
		Market:  market,
	}
	return f.Unsubscribe(msg, tickerChanName(market))
}
