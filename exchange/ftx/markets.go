package ftx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/arbitrage/exchange/ftx/data"
)

func (f FTX) TickerName(base string, quote string) (string, error) {
	jsonFile, err := os.Open("./exchange/ftx/data/markets.json")
	if err != nil {
		return "", fmt.Errorf("failed opening FTX markets file: %w", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", fmt.Errorf("failed reading FTX markets file: %w", err)
	}

	var markets data.Markets
	err = json.Unmarshal(byteValue, &markets)
	if err != nil {
		return "", fmt.Errorf("failed parsing FTX markets json: %w", err)
	}

	for _, symbol := range markets.Result {
		if symbol.Type == "spot" {
			if symbol.BaseCurrency == base && symbol.QuoteCurrency == quote {
				if symbol.Enabled {
					return symbol.Name, nil
				} else {
					return symbol.Name, fmt.Errorf("symbol disabled %s", symbol.Name)
				}
			}
		}
	}

	return "", fmt.Errorf("FTX not found: <%s-%s>", base, quote)
}
