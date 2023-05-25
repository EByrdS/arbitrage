package ftx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/arbitrage/ticker"
	"github.com/arbitrage/wsclient"
	"github.com/arbitrage/wsclient/lane"
)

func shouldIgnore(byteMessage []byte) bool {
	var received map[string]interface{}

	if err := json.Unmarshal(byteMessage, &received); err != nil {
		return false
	}

	return received["type"] == "subscribed" || received["type"] == "unsubscribed"
}

func discard(ignoreChan chan []byte) {
	for message := range ignoreChan {
		fmt.Printf("Discarding %s\n", message)
	}
}

func New() (FTX, error) {
	partner := FTX{
		Name:     "FTX",
		WSClient: wsclient.New(),
		Tickers:  map[string]chan ticker.Ticker{},
	}

	err := partner.Open("wss://ftx.com/ws")
	if err != nil {
		return partner, fmt.Errorf("error setting up ftx: %w", err)
	}

	ping, err := json.Marshal(map[string]string{"op": "ping"})
	if err != nil {
		return partner, fmt.Errorf("ping marshal error: %w", err)
	}
	partner.SetPing(ping, 10*time.Second)

	partner.Lanes["ignore"] = lane.New(shouldIgnore)
	go discard(partner.Lanes["ignore"].C)

	return partner, nil
}
