package connection

func New() *ExchangeConnection {
	return &ExchangeConnection{
		errorC:     make(chan error),
		bytesC:     make(chan []byte, 5),
		pingCloser: make(chan int),
	}
}
