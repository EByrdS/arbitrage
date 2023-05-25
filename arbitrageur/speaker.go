package arbitrageur

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/arbitrage/ticker"
)

func Speaker(
	increase float64,
) CompareFunc {
	fmt.Println("Building Speaker")

	return func(left ticker.Ticker, right ticker.Ticker) {

		leftAsk, err := strconv.ParseFloat(left.Ask, 32)
		if err != nil {
			fmt.Printf("Error converting %s Ask\n", left.Exchange)
			return
		}
		leftBid, err := strconv.ParseFloat(left.Bid, 32)
		if err != nil {
			fmt.Printf("Error converting %s Bid\n", left.Exchange)
			return
		}

		rightAsk, err := strconv.ParseFloat(right.Ask, 32)
		if err != nil {
			fmt.Printf("Error converting %s Ask\n", right.Exchange)
			return
		}
		rightBid, err := strconv.ParseFloat(right.Bid, 32)
		if err != nil {
			fmt.Printf("Error converting %s Bid\n", right.Exchange)
			return
		}

		if (leftBid > (rightAsk * (1 + increase))) || (rightBid > (leftAsk * (1 + increase))) {
			fmt.Printf("\n\n(%s) %s Ask: %s Bid: %s  %s~%s  Ask: %s Bid: %s %s (%s)\n\n",
				left.Identifier,
				left.Exchange,
				left.Ask,
				left.Bid,
				left.Market,
				right.Market,
				right.Ask,
				right.Bid,
				right.Exchange,
				right.Identifier,
			)

			message := fmt.Sprintf("Opportunity found between %s and %s "+
				" in the %s market", left.Exchange, right.Exchange,
				left.Market)
			cmd := exec.Command("say", "-v", "Victoria", message)

			err := cmd.Run()
			if err != nil {
				log.Printf("Error running command")
				return
			}
		}
	}
}
