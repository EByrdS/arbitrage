package arbitrageur

import (
	"fmt"

	"github.com/arbitrage/ticker"
)

func Printer() CompareFunc {
	logger.Println("Building Printer")

	return func(left ticker.Ticker, right ticker.Ticker) {

		if left.Identifier == "" || right.Identifier == "" {
			fmt.Print("\rWaiting for both tickers...")
			return
		}

		fmt.Printf("\r(%s - %s) Ask: %s Bid: %s  %s~%s  Ask: %s Bid: %s (%s - %s)",
			left.Exchange,
			left.Identifier,
			left.Ask,
			left.Bid,
			left.Market,
			right.Market,
			right.Ask,
			right.Bid,
			right.Identifier,
			right.Exchange,
		)
	}
}
