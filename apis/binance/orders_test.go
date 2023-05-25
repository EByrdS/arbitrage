package binance_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/arbitrage/apis/binance"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type CalledWith struct {
	function string
	params   map[string]interface{}
}

type RestMock struct {
	callStack      []CalledWith
	failFunc       map[string]bool
	stringResponse string
}

func (m *RestMock) Delete(endpoint string, payload io.Reader, securityLevel string) (
	*http.Response, error,
) {
	m.callStack = append(m.callStack, CalledWith{
		function: "Delete",
		params: map[string]interface{}{
			"endpoint":      endpoint,
			"payload":       payload,
			"securityLevel": securityLevel,
		},
	})
	if m.failFunc["Delete"] {
		return nil, fmt.Errorf("test error")
	}
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(m.stringResponse)),
	}, nil
}

func (m *RestMock) Get(endpoint string, params url.Values, securityLevel string) (
	*http.Response, error,
) {
	m.callStack = append(m.callStack, CalledWith{
		function: "Get",
		params: map[string]interface{}{
			"endpoint":      endpoint,
			"params":        params,
			"securityLevel": securityLevel,
		},
	})
	if m.failFunc["Get"] {
		return nil, fmt.Errorf("test error")
	}
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(m.stringResponse)),
	}, nil
}

func (m *RestMock) Post(endpoint string, body io.Reader, securityLevel string) (
	*http.Response, error,
) {
	m.callStack = append(m.callStack, CalledWith{
		function: "Post",
		params: map[string]interface{}{
			"endpoint":      endpoint,
			"body":          body,
			"securityLevel": securityLevel,
		},
	})
	if m.failFunc["Post"] {
		return nil, fmt.Errorf("test error")
	}
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(m.stringResponse)),
	}, nil
}

var _ = Describe("Orders", func() {
	Describe("NewOrder", func() {
		restMock := &RestMock{}
		restMock.stringResponse = "{\"symbol\":\"BTC-DOGE\",\"orderId\":30}"
		subject := binance.Binance{restMock}

		It("Calls endpoint with right parameters", func() {
			orderRequest := binance.OrderRequest{
				Symbol: "some-symbol",
				Side:   "some-side",
			}
			resp, err := subject.NewOrder(orderRequest)
			Ω(err).Should(BeNil())
			Ω(resp).Should(Equal(binance.OrderAPIResponse{
				Symbol:  "BTC-DOGE",
				OrderId: 30,
			}))

			Ω(len(restMock.callStack)).Should(Equal(1))

			expectedBody, err := json.Marshal(orderRequest)
			Ω(err).Should(BeNil())
			Ω(restMock.callStack[0]).Should(Equal(CalledWith{
				function: "Post",
				params: map[string]interface{}{
					"endpoint":      "/api/v3/order",
					"body":          bytes.NewReader(expectedBody),
					"securityLevel": "TRADE",
				},
			}))
		})
	})

	Describe("CancelOrder", func() {
		restMock := &RestMock{}
		restMock.stringResponse = "{\"symbol\":\"BTC-DOGE\",\"orderId\":10}"
		subject := binance.Binance{restMock}

		It("Calls endpoint with right parameters", func() {
			resp, err := subject.CancelOrder("BTC-DOGE", 10)
			Ω(resp).Should(Equal(binance.CancelAPIResponse{
				Symbol:  "BTC-DOGE",
				OrderId: 10,
			}))
			Ω(err).Should(BeNil())

			Ω(len(restMock.callStack)).Should(Equal(1))

			expectedData := []byte("{\"orderId\":\"10\",\"symbol\":\"BTC-DOGE\"}")
			Ω(err).Should(BeNil())
			Ω(restMock.callStack[0]).Should(Equal(CalledWith{
				function: "Delete",
				params: map[string]interface{}{
					"endpoint":      "/api/v3/order",
					"payload":       bytes.NewReader(expectedData),
					"securityLevel": "TRADE",
				},
			}))
		})

	})
})
