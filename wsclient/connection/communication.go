package connection

import (
	"fmt"
	"strings"

	"time"

	"github.com/gorilla/websocket"
)

func (middleman ExchangeConnection) SendJSON(messageJSON interface{}) error {
	err := middleman.websocketConn.WriteJSON(messageJSON)
	if err != nil {
		return fmt.Errorf("Sending error: %v", err)
	}
	return nil
}

func stream(middleman *ExchangeConnection) {
	defer middleman.websocketConn.Close()
	defer close(middleman.BytesC())

	for {
		_, message, err := middleman.websocketConn.ReadMessage()
		if err != nil {
			if c, k := err.(*websocket.CloseError); k && (c.Code == 1000) {
				break
			}
			middleman.ErrorC() <- fmt.Errorf("read err: %v", err)
			continue
		}

		if strings.Contains(string(message), "error") {
			middleman.ErrorC() <- fmt.Errorf("Message with error %s", string(message))
			continue
		}

		middleman.BytesC() <- message
	}
}

func (middleman *ExchangeConnection) Open(url string) error {
	openConnection, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("Open error %v", err)
	}

	middleman.websocketConn = openConnection

	go stream(middleman)

	return nil
}

func (middleman ExchangeConnection) Close(message string) error {
	closeMessage := websocket.
		FormatCloseMessage(websocket.CloseNormalClosure, message)
	err := middleman.websocketConn.WriteMessage(websocket.CloseMessage, closeMessage)
	if err != nil {
		return fmt.Errorf("websocket close err: %v", err)
	}
	close(middleman.pingCloser)
	return nil
}

func (middleman ExchangeConnection) SetPing(ping []byte, period time.Duration) {
	go func() {
		for {
			select {
			case <-time.After(period):

				// Pongs handled by (*websocket.Conn).SetPongHandler
				// Default behavior will drop the message
				err := middleman.websocketConn.WriteMessage(websocket.PingMessage, ping)
				if err != nil {
					middleman.ErrorC() <- fmt.Errorf("PeriodicPing error: %v", err)
				}
			case <-middleman.pingCloser:
				return
			}
		}
	}()
}
