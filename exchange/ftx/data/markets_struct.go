package data

type Markets struct {
	Success bool     `json:"success"`
	Result  []Symbol `json:"result"`
}

type Symbol struct {
	Name                  string  `json:"name"`
	Enabled               bool    `json:"enabled"`
	PostOnly              bool    `json:"postOnly"`
	PriceIncrement        float64 `json:"priceIncrement"`
	SizeIncrement         float64 `json:"sizeIncrement"`
	MinProvideSize        float64 `json:"minProvideSize"`
	Last                  float64 `json:"last"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	Price                 float64 `json:"price"`
	Type                  string  `json:"type"` // predefined types?
	BaseCurrency          string  `json:"baseCurrency"`
	QuoteCurrency         string  `json:"quoteCurrency"`
	Underlying            string  `json:"underlying"`
	Restricted            bool    `json:"restricted"`
	HighLeverageFeeExempt bool    `json:"highLeverageFeeExempt"`
	Change1h              float64 `json:"change1h"`
	Change24h             float64 `json:"change24h"`
	ChangeBod             float64 `json:"changeBod"`
	QuoteVolume24h        float64 `json:"quoteVolume24h"`
	VolumeUsd24h          float64 `json:"volumeUsd24h"`
}
