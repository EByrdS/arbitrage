package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/arbitrage/apis/ftx/authentication"
)

func Get(endpoint string, params url.Values) (*http.Response, []byte, error) {
	ftxURL := apiURL()
	ftxURL.Path = "/api" + endpoint
	ftxURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", ftxURL.String(), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating GET request: %w", err)
	}

	authentication.AddHeaders(req)

	// req.Header.Add("FTX-SUBACCOUNT") // if using subaccount

	resp, err := client.Do(req)
	if err != nil {
		return resp, nil, fmt.Errorf("error executing GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, body, err
	}

	result, err := substractApiResult(body)

	return resp, result, err
}
