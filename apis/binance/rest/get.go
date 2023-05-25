package rest

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/arbitrage/apis/binance/authentication"
)

func (rest Rest) Get(endpoint string, params url.Values, securityLevel string) (*http.Response, error) {
	binanceURL := apiURL()
	binanceURL.Path = endpoint
	binanceURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", binanceURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating GET request: %w", err)
	}

	authentication.AddHeaders(req, securityLevel)

	fmt.Printf("RequestURI: %v\n", req.URL.RequestURI())
	resp, err := client.Do(req) // will need to close the connection from the caller
	if err != nil {
		return resp, fmt.Errorf("error executing GET request: %w", err)
	}

	return resp, err
}
