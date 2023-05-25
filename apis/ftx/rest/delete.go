package rest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/arbitrage/apis/ftx/authentication"
)

func Delete(endpoint string, payload io.Reader) (*http.Response, error) {
	ftxURL := apiURL()
	ftxURL.Path = "/api" + endpoint

	req, err := http.NewRequest("DELETE", ftxURL.String(), payload)
	if err != nil {
		return nil, fmt.Errorf("failed creating DELETE request: %w", err)
	}

	authentication.AddHeaders(req)

	// req.Header.Add("FTX-SUBACCOUNT") // if using subaccount

	resp, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("error executing DELETE request: %w", err)
	}

	return resp, err
}
