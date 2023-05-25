package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/arbitrage/exchange/binance/data"
)

func (b Binance) TickerName(base string, quote string) (string, error) {
	jsonFile, err := os.Open("./exchange/binance/data/markets.json")
	if err != nil {
		return "", fmt.Errorf("failed opening Binance markets file: %w", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", fmt.Errorf("failed reading Binance markets file: %w", err)
	}

	var markets data.Markets
	err = json.Unmarshal(byteValue, &markets)
	if err != nil {
		return "", fmt.Errorf("failed parsing Binance markets json: %w", err)
	}

	for _, symbol := range markets.Symbols {
		if symbol.BaseAsset == base && symbol.QuoteAsset == quote {
			tickerName := strings.ToLower(symbol.BaseAsset) +
				strings.ToLower(symbol.QuoteAsset) + "@bookTicker"
			return tickerName, nil
		}
	}

	return "", fmt.Errorf("Binance not found: <%s-%s>", base, quote)
}
