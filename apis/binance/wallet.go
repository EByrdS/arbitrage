package binance

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type AddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

func (api Binance) GetDepositAddress(coin string, network string,
) (addressResponse AddressResponse, err error) {
	params := url.Values{}

	params.Add("coin", coin)
	if network != "" {
		params.Add("network", network)
	}

	resp, err := api.Rest.Get(
		"/sapi/v1/capital/deposit/address",
		params,
		"USER_DATA",
	)
	if err != nil {
		return addressResponse, fmt.
			Errorf("error getting deposit address: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&addressResponse)
	if err != nil {
		return addressResponse, fmt.
			Errorf("error decoding get deposit address response: %w", err)
	}

	return addressResponse, nil
}
