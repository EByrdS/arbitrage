package wsclient

import "fmt"

func AssignChannels(c *WSClient) {
Out:
	for message := range c.BytesC() {
		for _, lane := range c.Lanes {
			if lane.Belongs(message) {
				lane.C <- message
				continue Out
			}
		}

		c.ErrorC() <- fmt.Errorf("message without lane (out of %d): %s", len(c.Lanes), string(message))
	}
}
