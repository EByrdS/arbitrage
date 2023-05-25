package connection

func (middleman ExchangeConnection) ErrorC() chan error {
	return middleman.errorC
}

func (middleman ExchangeConnection) BytesC() chan []byte {
	return middleman.bytesC
}
