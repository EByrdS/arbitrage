package ftx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/arbitrage/apis/ftx/rest"
)

// https://docs.ftx.com/#get-balances

type AddressResponse struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Method  string `json:"method"`
	Coin    string `json:"coin"`
}

type WithdrawalResponse struct {
	Coin    string  `json:"coin"`
	Address string  `json:"address"`
	Tag     string  `json:"tag"`
	Fee     float64 `json:"fee"`
	Id      int     `json:"id"`
	Size    string  `json:"size"`
	Status  string  `json:"status"`
	Time    string  `json:"time"`
	TxId    string  `json:"txid"`
}

type BalancesResponse map[string][]BalanceResponse

type BalanceResponse struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	SpotBorrow             float64 `json:"spotBorrow"`
	Total                  float64 `json:"total"`
	USDValue               float64 `json:"usdValue"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}

func (api FTX) GetDepositAddress(coin string, network string,
) (addressResponse AddressResponse, err error) {
	params := url.Values{}

	// for coins available on different blockchains
	if network != "" {
		params.Add("method", network) // optional
	}

	_, result, err := rest.Get(
		fmt.Sprintf("/wallet/deposit_address/%s", coin),
		params)
	if err != nil {
		return addressResponse, fmt.
			Errorf("error getting deposit address: %w", err)
	}

	if err := json.Unmarshal(result, &addressResponse); err != nil {
		return addressResponse, fmt.Errorf("failed unmarshalling result")
	}

	return addressResponse, nil
}

func (api FTX) RequestWithdrawal(
	coin string, size float64, address string, tag string, method string,
	password string, code string,
) (withdrawalResponse WithdrawalResponse, err error) {
	payload := make(map[string]interface{})
	payload["coin"] = coin
	payload["size"] = size
	payload["address"] = address
	if tag != "" {
		payload["tag"] = tag
	}
	if method != "" {
		payload["method"] = method
	}
	if password != "" {
		payload["password"] = password
	}
	if code != "" {
		payload["code"] = code
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return
	}

	_, result, err := rest.Post("/wallet/withdrawals", bytes.NewReader(data))
	if err != nil {
		return withdrawalResponse, fmt.
			Errorf("error posting to /wallet/withdrawals: %w", err)
	}

	if err = json.Unmarshal(result, &withdrawalResponse); err != nil {
		return withdrawalResponse, fmt.
			Errorf("failed unmarshalling response: %w", err)
	}

	return withdrawalResponse, nil
}

func (api FTX) GetAllBalances() (balancesResponse BalancesResponse, err error) {
	_, result, err := rest.Get("/wallet/all_balances", url.Values{})
	if err != nil {
		return
	}

	if err := json.Unmarshal(result, &balancesResponse); err != nil {
		fmt.Printf("%v\n", result)
		return balancesResponse, fmt.
			Errorf("failed unmarshalling response: %w", err)
	}

	return balancesResponse, nil
}
