package ftx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/arbitrage/apis/ftx/rest"
)

type SingleOrderResponse struct {
	Success bool          `json:"success"`
	Result  OrderResponse `json:"result"`
}

type MultipleOrderResponse struct {
	Success bool            `json:"success"`
	Result  []OrderResponse `json:"result"`
}

type StringAPIResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

type OrderResponse struct {
	Market        string    `json:"market"`     // e.g. "BTC/USD" for spot, "XRP-PERP" for futures
	Price         float64   `json:"price"`      // Send null for market orders
	Side          string    `json:"side"`       // "buy" or "sell"
	Size          float64   `json:"size"`       // :nodoc:
	Type          string    `json:"type"`       // "limit" or "market"
	ReduceOnly    bool      `json:"reduceOnly"` // optional
	IOC           bool      `json:"ioc"`        // optional
	PostOnly      bool      `json:"postOnly"`   // optional
	ClientId      string    `json:"clientId"`   // optional; client order id
	Id            int64     `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	FilledSize    float64   `json:"filledSize"`
	RemainingSize float64   `json:"remainingSize"`
	Status        string    `json:"status"`
}

type OrderRequest struct {
	Market            string  `json:"market"`            // e.g. "BTC/USD" for spot, "XRP-PERP" for futures
	Price             float64 `json:"price"`             // Send null for market orders
	Side              string  `json:"side"`              // "buy" or "sell"
	Size              float64 `json:"size"`              // :nodoc:
	Type              string  `json:"type"`              // "limit" or "market"
	ReduceOnly        bool    `json:"reduceOnly"`        // optional
	IOC               bool    `json:"ioc"`               // optional
	PostOnly          bool    `json:"postOnly"`          // optional
	ClientId          string  `json:"clientId"`          // optional; client order id
	RejectOnPriceBand bool    `json:"rejectOnPriceBand"` // optional; if the order should be rejected if its price would instead be adjusted due to price bands
	RejectAfterTs     int64   `json:"rejectAfterTs"`     // optional; if the order would be put into the placement queue after this timestamp, instead reject it.
}

func (api FTX) PlaceOrder(orderRequest OrderRequest) (orderResponse OrderResponse, err error) {
	data, err := json.Marshal(orderRequest)
	if err != nil {
		return orderResponse, fmt.
			Errorf("failed marshaling data to place order: %w", err)
	}

	_, result, err := rest.Post("/orders", bytes.NewReader(data))
	if err != nil {
		return orderResponse, fmt.Errorf("error posting to /orders: %w", err)
	}

	err = json.Unmarshal(result, &orderResponse)
	if err != nil {
		return orderResponse, fmt.
			Errorf("error decoding place order response: %w", err)
	}

	return orderResponse, err
}

func (api FTX) GetOpenOrders(market string) (orderResponses []OrderResponse, err error) {
	params := url.Values{}
	params.Add("market", market)

	_, result, err := rest.Get("/orders", params)
	if err != nil {
		return orderResponses, fmt.Errorf("error getting open orders: %w", err)
	}

	multipleOrderResponse := new(MultipleOrderResponse)
	err = json.Unmarshal(result, &multipleOrderResponse)
	if err != nil {
		return orderResponses, fmt.
			Errorf("error decoding get open orders response: %w", err)
	}

	return multipleOrderResponse.Result, nil
}

func (api FTX) CancelOrder(orderId int64) (string, error) {
	resp, err := rest.Delete(fmt.Sprintf("/orders/%v", orderId), nil)
	if err != nil {
		return "", fmt.Errorf("error cancelling order: %w", err)
	}
	defer resp.Body.Close()

	stringAPIResponse := new(StringAPIResponse)
	err = json.NewDecoder(resp.Body).Decode(&stringAPIResponse)
	if err != nil {
		return "", fmt.Errorf("error decoding cancel order response: %w", err)
	}

	return stringAPIResponse.Result, nil
}

func (api FTX) CancelAllOrders(market string, side string, conditionalsOnly bool, limitOnly bool) (string, error) {
	payload := make(map[string]string)
	if market != "" {
		payload["market"] = market
	}
	if side != "" {
		payload["side"] = side
	}
	if conditionalsOnly {
		payload["conditionalOrdersOnly"] = "true"
	}
	if limitOnly {
		payload["limitOrdersOnly"] = "true"
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed marshaling data to place order: %w", err)
	}

	resp, err := rest.Delete("/orders", bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("error cancelling all orders: %w", err)
	}
	defer resp.Body.Close()

	stringAPIResponse := new(StringAPIResponse)
	err = json.NewDecoder(resp.Body).Decode(&stringAPIResponse)
	if err != nil {
		return "", fmt.
			Errorf("error decoding cancel all orders response: %w", err)
	}

	return stringAPIResponse.Result, nil
}
