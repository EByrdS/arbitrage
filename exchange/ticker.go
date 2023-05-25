package exchange

import (
	"fmt"

	"github.com/arbitrage/ticker"
)

func GetTicker(
	partner Partner,
	base string,
	quote string,
) (
	chan ticker.Ticker,
	chan int,
	error) {

	var tickerChan chan ticker.Ticker
	var stopper chan int
	var err error

	market, err := partner.TickerName(base, quote)
	if err != nil {
		fmt.Printf("Market name error %s-%s: %+v", base, quote, err)
		return tickerChan, stopper, fmt.Errorf("market name error %s-%s, %+v", base, quote, err)
	}

	tickerChan, err = partner.GetTicker(market)
	if err != nil {
		fmt.Println("subscribe error:", err)
		return tickerChan, stopper, fmt.Errorf("Subscription error: %v", err)
	}

	go func() {
		<-stopper

		fmt.Printf("interrupting Ticker %s\n", market)

		err := partner.StopTicker(market)
		if err != nil {
			// Lane channel might still be open!
			fmt.Println("unsubscribe err:", err)
			partner.ErrorC() <- fmt.Errorf("Ticker stop err: %v", err)
		}
	}()

	stopper = make(chan int)

	return tickerChan, stopper, nil
}
