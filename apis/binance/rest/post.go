package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/arbitrage/apis/binance/authentication"
)

func (rest Rest) Post(endpoint string, body io.Reader, securityLevel string) (
	*http.Response, error) {
	binanceURL := apiURL()
	binanceURL.Path = endpoint

	req, err := http.NewRequest("POST", binanceURL.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed creating POST request: %w", err)
	}

	authentication.AddHeaders(req, securityLevel)

	resp, err := client.Do(req) // will need to close the reso connection from the caller
	if err != nil {
		return resp, fmt.Errorf("error executing POST request: %w", err)
	}

	return resp, err
}
