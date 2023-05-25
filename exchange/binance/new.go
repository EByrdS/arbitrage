package binance

import (
	"encoding/json"
	"fmt"

	"github.com/arbitrage/ticker"
	"github.com/arbitrage/wsclient"
	"github.com/arbitrage/wsclient/lane"
)

func shouldIgnore(byteMessage []byte) bool {
	var received map[string]interface{}

	if err := json.Unmarshal(byteMessage, &received); err != nil {
		return false
	}

	value, ok := received["result"]
	return ok && value == nil
}

func discard(ignoreChan chan []byte) {
	for message := range ignoreChan {
		fmt.Printf("Discarding %s\n", message)
	}
}

func New() (Binance, error) {
	partner := Binance{
		Name:     "Binance",
		WSClient: wsclient.New(),
		Tickers:  map[string]chan ticker.Ticker{},
	}

	err := partner.Open("wss://stream.binance.com:9443/ws")
	if err != nil {
		return partner, fmt.Errorf("error setting up binance: %w", err)
	}

	partner.Lanes["ignore"] = lane.New(shouldIgnore)
	go discard(partner.Lanes["ignore"].C)

	return partner, nil
}
