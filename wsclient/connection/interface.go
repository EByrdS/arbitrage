package connection

import "time"

type Connector interface {
	Open(url string) error
	Close(message string) error
	SendJSON(messageJSON interface{}) error
	SetPing(ping []byte, period time.Duration)

	BytesC() chan []byte
	ErrorC() chan error
}
