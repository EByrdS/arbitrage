package binance

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type OrderRequest struct {
	Symbol           string `json:"symbol"`
	Side             string `json:"side"`
	Type             string `json:"type"`
	TimeInForce      string `json:"timeInForce"`      // optional
	Quantity         string `json:"quantity"`         // optional
	QuoteOrderQty    string `json:"quoteOrderQty"`    // optional
	Price            string `json:"price"`            // optional
	NewClientOrderId string `json:"newClientOrderId"` // optional
	StopPrice        string `json:"stopPrice"`        // optional
	IcebergQty       string `json:"icebergQty"`       // optional
	NewOrderRespType string `json:"newOrderRespType"` // optional
}

type OrderAPIResponse struct {
	Symbol              string `json:"symbol"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	Fills               []Fill `json:"fills"`
}

type Fill struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeId         int    `json:"tradeId"`
}

func (api Binance) NewOrder(orderRequest OrderRequest) (
	orderResponse OrderAPIResponse, err error,
) {
	data, err := json.Marshal(orderRequest)
	if err != nil {
		return orderResponse, fmt.
			Errorf("failed marhsaling data to place order: %w", err)
	}

	resp, err := api.Rest.Post("/api/v3/order", bytes.NewReader(data), "TRADE")
	if err != nil {
		return orderResponse, fmt.
			Errorf("error posting to /api/v3/order: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&orderResponse)
	if err != nil {
		return orderResponse, fmt.
			Errorf("error decoding new order response: %w", err)
	}

	return orderResponse, err
}

type CancelAPIResponse struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

func (api Binance) CancelOrder(symbol string, orderId int) (
	cancelAPIResponse CancelAPIResponse, err error,
) {
	payload := make(map[string]string)
	payload["symbol"] = symbol
	payload["orderId"] = fmt.Sprint(orderId)

	data, err := json.Marshal(payload)
	if err != nil {
		return cancelAPIResponse, fmt.Errorf("failed marshaling data to cancel order: %w", err)
	}

	resp, err := api.Rest.Delete("/api/v3/order", bytes.NewReader(data), "TRADE")
	if err != nil {
		return cancelAPIResponse, fmt.Errorf("error cancelling order: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cancelAPIResponse)
	if err != nil {
		return cancelAPIResponse, fmt.
			Errorf("Error decoding cancel order response: %w", err)
	}

	return cancelAPIResponse, nil
}
