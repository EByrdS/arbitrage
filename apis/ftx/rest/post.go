package rest

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/arbitrage/apis/ftx/authentication"
)

func Post(endpoint string, payload io.Reader) (*http.Response, []byte, error) {
	ftxURL := apiURL()
	ftxURL.Path = "/api" + endpoint

	req, err := http.NewRequest("POST", ftxURL.String(), payload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating POST request: %w", err)
	}

	authentication.AddHeaders(req)

	// req.Header.Add("FTX-SUBACCOUNT") // if using subaccount

	resp, err := client.Do(req) // will need to close the resp connection from the caller
	if err != nil {
		return resp, nil, fmt.Errorf("error executing POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, body, err
	}

	result, err := substractApiResult(body)

	return resp, result, err
}
