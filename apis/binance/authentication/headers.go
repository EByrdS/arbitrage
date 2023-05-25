package authentication

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func AddHeaders(req *http.Request, securityLevel string) {
	// TODO: Move apiKey and apiSecret to struct creation
	apiKey, exists := os.LookupEnv("BINANCE_API_KEY")
	if !exists {
		panic("No env variable BINANCE_API_KEY")
	}

	apiSecret, exists := os.LookupEnv("BINANCE_API_SECRET")
	if !exists {
		panic("No env variable BINANCE_API_SECRET")
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	q := req.URL.Query()
	q.Add("timestamp", fmt.Sprint(timestamp)) // check that it is set in POST

	sendKey, sign := SecurityLevel(securityLevel)
	fmt.Printf("sendKey: %t, sign: %t\n", sendKey, sign)
	if sendKey {
		req.Header.Add("X-MBX-APIKEY", apiKey)
		if sign {
			q.Add("signature", Sign(req.URL.RawQuery, timestamp, apiSecret))
		}
	}

	req.URL.RawQuery = q.Encode()

	if req.Method == "GET" {
		req.Header.Add("Content-Type", "application/json") // ? or always the other one?
	} else if req.Method == "POST" || req.Method == "PUT" || req.Method == "DELETE" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	fmt.Printf("Headers: %v\n", req.Header)
	fmt.Printf("Query: %v\n", req.URL.Query())
}
