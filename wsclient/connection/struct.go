package connection

import "github.com/gorilla/websocket"

type ExchangeConnection struct {
	errorC     chan error
	bytesC     chan []byte
	pingCloser chan int

	websocketConn *websocket.Conn
}
