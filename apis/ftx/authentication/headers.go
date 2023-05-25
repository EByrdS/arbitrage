package authentication

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func AddHeaders(req *http.Request) {

	// TODO: Move apiKey and apiSecret to struct creation
	apiKey, exists := os.LookupEnv("FTX_API_KEY")
	if !exists {
		panic("No environment variable FTX_API_KEY")
	}
	apiSecret, exists := os.LookupEnv("FTX_API_SECRET")
	if !exists {
		panic("No environment variable FTX_API_SECRET")
	}

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// check timestamp against https://otc.ftx.com/api/time

	buffer := new(strings.Builder)
	if req.Body != nil {
		if _, err := io.Copy(buffer, req.Body); err != nil {
			panic(fmt.Errorf("cannot convert request body to string: %w", err))
		}
	}

	req.Header.Add("FTX-KEY", apiKey)
	req.Header.Add("FTX-TS", fmt.Sprint(timestamp))
	req.Header.Add("FTX-SIGN",
		Sign(req.Method, req.URL.Path, buffer.String(), timestamp, apiSecret),
	)
	req.Header.Add("Content-Type", "application/json")
}
