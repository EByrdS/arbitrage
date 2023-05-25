package wsclient

import (
	"fmt"

	"github.com/arbitrage/wsclient/lane"
)

// An exchange can subscribe to multiple things
// it is usefull as a separate function to handle channels independently
func (c *WSClient) Subscribe(
	message interface{},
	laneName string,
	scanner func([]byte) bool,
) (lane.Lane, error) {
	// fmt.Printf("Subscribing to %s, with %+v\n", laneName, message)

	var newLane lane.Lane

	if _, exists := c.Lanes[laneName]; exists {
		return newLane, fmt.Errorf("lane already exists: %s", laneName)
	}

	err := c.SendJSON(message)
	if err != nil {
		return newLane, fmt.Errorf("error in SendJSON <%+v>: %v", message, err)
	}

	newLane = lane.New(scanner)
	c.Lanes[laneName] = newLane

	return newLane, nil
}

// You can unsubscribe from multiple things
// usefull as abstracted function to handle channels independently
func (c *WSClient) Unsubscribe(
	message interface{},
	laneName string,
) error {
	err := c.SendJSON(message)
	if err != nil {
		return fmt.Errorf("unsubscription error <%+v>: %v", message, err)
	}

	if lane, exists := c.Lanes[laneName]; exists {
		close(lane.C)
		delete(c.Lanes, laneName)
	} else {
		return fmt.Errorf("lane not found: %s", laneName)
	}

	return nil
}
