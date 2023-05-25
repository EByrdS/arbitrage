package binance

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arbitrage/ticker"
)

func tickerChanName(market string) string {
	return "ticker-" + market
}

func tickerBelongs(market string) func([]byte) bool {
	return func(byteMessage []byte) bool {
		var received inTicker

		if err := json.Unmarshal(byteMessage, &received); err != nil {
			return false
		}

		return received.UpdateId != 0 && received.Symbol == strings.ToUpper(strings.Split(market, "@")[0])
	}
}

func (b Binance) GetTicker(market string) (chan ticker.Ticker, error) {
	var tickerChan chan ticker.Ticker
	msg := outMessage{
		Method: "SUBSCRIBE",
		Params: []string{market},
		Id:     1, // ?? check docs
	}

	channelName := tickerChanName(market)

	byteLane, err := b.Subscribe(msg, channelName, tickerBelongs(market))
	if err != nil {
		return tickerChan, fmt.Errorf("Subscribe error: %v", err)
	}

	destination := make(chan ticker.Ticker)
	b.Tickers[channelName] = destination
	go asTicker(byteLane.C, destination, b.ErrorC())

	return destination, nil
}

func (b Binance) StopTicker(market string) error {
	msg := outMessage{
		Method: "UNSUBSCRIBE",
		Params: []string{market},
		Id:     1, // ?? check docs
	}
	return b.Unsubscribe(msg, tickerChanName(market))
}
