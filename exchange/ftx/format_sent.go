package ftx

type outMessage struct {
	Op      string `json:"op"`
	Channel string `json:"channel"`
	Market  string `json:"market"`
}
