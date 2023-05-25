package rest

import (
	"encoding/json"
	"fmt"
)

type SuccessResultResponse struct {
	Success bool            `json:"success"`
	Result  json.RawMessage `json:"result"`
}

func substractApiResult(body []byte) ([]byte, error) {
	var apiResponse SuccessResultResponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed unmarhsalling response")
	}

	if !apiResponse.Success {
		return apiResponse.Result, fmt.Errorf("api call did not succeed")
	}

	return apiResponse.Result, nil
}
