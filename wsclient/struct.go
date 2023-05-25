package wsclient

import (
	"github.com/arbitrage/wsclient/connection"
	"github.com/arbitrage/wsclient/lane"
)

type WSClient struct {
	connection.Connector
	Lanes map[string]lane.Lane
}
