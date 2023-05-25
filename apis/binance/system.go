package binance

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type StatusResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (api Binance) GetSystemStatus() (statusResponse StatusResponse, err error) {
	resp, err := api.Rest.Get("/sapi/v1/system/status", url.Values{}, "NONE")
	if err != nil {
		return statusResponse, fmt.Errorf("error getting system status: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&statusResponse)
	if err != nil {
		return statusResponse, fmt.
			Errorf("error decoding get system status response: %w", err)
	}

	return statusResponse, nil
}
