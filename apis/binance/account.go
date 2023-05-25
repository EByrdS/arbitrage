package binance

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInformation struct {
	MakerCommission  int       `json:"makerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	BuyerCommission  int       `json:"buyerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       int       `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
	Permissions      []string  `json:"permissions"`
}

func (api Binance) GetAccountInformation() (accountInformation AccountInformation, err error) {
	resp, err := api.Rest.Get(
		"/api/v3/account",
		url.Values{},
		"USER_DATA",
	)
	if err != nil {
		return accountInformation, fmt.
			Errorf("error getting account information: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&accountInformation)
	if err != nil {
		return accountInformation, fmt.
			Errorf("error decoding deposit address response: %w", err)
	}

	return accountInformation, nil
}

type AccountStatus struct {
	Data string `json:"data"`
}

func (api Binance) GetAccountStatus() (accountStatus AccountStatus, err error) {
	resp, err := api.Rest.Get(
		"/sapi/v1/account/status",
		url.Values{},
		"USER_DATA",
	)
	if err != nil {
		return accountStatus, fmt.
			Errorf("error getting account status: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&accountStatus)
	if err != nil {
		return accountStatus, fmt.
			Errorf("error decoding account statis: %w", err)
	}

	return accountStatus, nil
}
