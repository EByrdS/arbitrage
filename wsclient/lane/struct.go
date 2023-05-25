package lane

type Lane struct {
	C       chan []byte
	Belongs func([]byte) bool
}
