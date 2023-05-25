package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/arbitrage/apis/binance/authentication"
)

func (rest Rest) Delete(endpoint string, payload io.Reader, securityLevel string) (
	*http.Response, error,
) {
	binanceURL := apiURL()
	binanceURL.Path = endpoint

	req, err := http.NewRequest("DELETE", binanceURL.String(), payload)
	if err != nil {
		return nil, fmt.Errorf("failed creating DELETE request: %w", err)
	}

	authentication.AddHeaders(req, securityLevel)

	resp, err := client.Do(req) // will need to close the resp connection from the caller
	if err != nil {
		return resp, fmt.Errorf("error executing DELETE request: %w", err)
	}

	return resp, err
}
