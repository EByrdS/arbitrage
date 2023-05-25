package wsclient

import (
	"github.com/arbitrage/wsclient/connection"
	"github.com/arbitrage/wsclient/lane"
)

func New() *WSClient {
	c := &WSClient{
		Connector: connection.New(),
		Lanes:     map[string]lane.Lane{},
	}

	go AssignChannels(c)

	return c
}
