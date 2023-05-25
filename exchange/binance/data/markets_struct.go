package data

type Markets struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int               `json:"serverTime"`
	RateLimits      []RateLimits      `json:"rateLimits"`
	ExchangeFilters []ExchangeFilters `json:"exchangeFilters"`
	Symbols         []Symbols         `json:"symbols"`
}

type RateLimits struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

type ExchangeFilters struct {
}

type Symbols struct {
	Symbol                     string       `json:"symbol"`
	Status                     string       `json:"status"`
	BaseAsset                  string       `json:"baseAsset"`
	BaseAssetPrecision         int          `json:"baseAssetPrecision"`
	QuoteAsset                 string       `json:"quoteAsset"`
	QuotePrecision             int          `json:"quotePrecision"`
	QuoteAssetPrecision        int          `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int          `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int          `json:"quoteCommissionPrecision"`
	OrderTypes                 []OrderType  `json:"orderTypes"`
	IcebergAllowed             bool         `json:"icebergAllowed"`
	OcoAllowed                 bool         `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool         `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool         `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool         `json:"isMarginTradingAllowed"`
	Filters                    []Filter     `json:"filters"`
	Permissions                []Permission `json:"permissions"`
}

type OrderType string

const (
	Limit           OrderType = "LIMIT"
	LimitMaker      OrderType = "LIMIT_MAKER"
	Market          OrderType = "Market"
	StopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	TakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
)

type Filter struct {
	FilterType       FilterType `json:"filterType"`
	MinPrice         *string    `json:"minPrice"`
	MaxPrice         *string    `json:"maxPrize"`
	TickSize         *string    `json:"tickSize"`
	MultiplierUp     *string    `json:"multiplierUp"`
	MultiplierDown   *string    `json:"multiplierDown"`
	AvgPriceMins     *int       `json:"avgPriceMins"`
	MinQty           *string    `json:"MinQty"`
	MaxQty           *string    `json:"MaxQty"`
	StepSize         *string    `json:"stepSize"`
	MinNotional      *string    `json:"minNotional"`
	ApplyToMarket    *bool      `json:"applyToMarket"`
	Limit            *int       `json:"limit"`
	MaxNumOrders     *int       `json:"maxNumOrders"`
	MaxNumAlgoOrders *int       `json:"maxNumAlgoOrders"`
}

type FilterType string

const (
	PriceFilter      FilterType = "PRICE_FILTER"
	PercentPrice     FilterType = "PERCENT_PRICE"
	LotSize          FilterType = "LOT_SIZE"
	MinNotional      FilterType = "MIN_NOTIONAL"
	IcebergParts     FilterType = "ICEBERG_PARTS"
	MarketLotSize    FilterType = "MARKET_LOT_SIZE"
	MaxNumOrders     FilterType = "MAX_NUM_ORDERS"
	MaxNumAlgoOrders FilterType = "MAX_NUM_ALGO_ORDERS"
)

type Permission string

const (
	Spot   Permission = "SPOT"
	Margin Permission = "MARGIN"
)
