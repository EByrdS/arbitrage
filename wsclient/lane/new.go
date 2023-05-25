package lane

func New(
	verifier func([]byte) bool,
) Lane {
	return Lane{
		C:       make(chan []byte, 5),
		Belongs: verifier,
	}
}
