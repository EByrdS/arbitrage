package main

import (
	"fmt"
	logPkg "log"
	"os"
	"os/signal"
	"time"

	"github.com/arbitrage/arbitrageur"
	"github.com/arbitrage/exchange"
	"github.com/arbitrage/exchange/binance"
	"github.com/arbitrage/exchange/ftx"
)

func main() {
	var log = logPkg.New(os.Stdout, "Arbitrage) ", 0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// FTX
	partnerFTX, err := ftx.New()
	if err != nil {
		log.Fatal("Error creating FTX:", err)
		return
	}

	// "DOGE/BTC"
	dogebtcFTXC, dogebtcFTXI, err := exchange.GetTicker(partnerFTX, "DOGE", "BTC")
	if err != nil {
		log.Fatal("Error getting FTX doge-btc ticker:", err)
		return
	}

	// LTC/BTC
	ltcbtcFTXC, ltcbtcFTXI, err := exchange.GetTicker(partnerFTX, "LTC", "BTC")
	if err != nil {
		log.Fatal("Error getting FTX ltc-btc ticker:", err)
		return
	}

	// Binance
	partnerBinance, err := binance.New()
	if err != nil {
		log.Fatal("Error creating Binance:", err)
		return
	}

	// "dogebtc@bookTicker"
	dogebtcBinanceC, dogebtcBinanceI, err := exchange.GetTicker(partnerBinance, "DOGE", "BTC")
	if err != nil {
		log.Fatal("Error getting Binance doge-btc ticker:", err)
		return
	}

	ltcbtcBinanceC, ltcbtcBinanceI, err := exchange.GetTicker(partnerBinance, "LTC", "BTC")
	if err != nil {
		log.Fatal("Error getting Binance ltc-btc ticker:", err)
		return
	}

	dogebtcInterrupter := make(chan int)
	go arbitrageur.Compare(
		dogebtcFTXC,
		dogebtcBinanceC,
		arbitrageur.Printer(),
		dogebtcInterrupter,
	)

	ltcbtcInterrupter := make(chan int)
	go arbitrageur.Compare(
		ltcbtcFTXC,
		ltcbtcBinanceC,
		arbitrageur.Printer(),
		ltcbtcInterrupter,
	)

	closeAll := func() {
		fmt.Println("Interrupting tickers...")

		close(dogebtcBinanceI)
		close(dogebtcFTXI)
		close(dogebtcInterrupter)

		<-time.After(time.Second)

		close(ltcbtcBinanceI)
		close(ltcbtcFTXI)
		close(ltcbtcInterrupter)

		<-time.After(time.Second)
		fmt.Println("Stopping exchanges...")

		if err := partnerFTX.Close("Thanks!"); err != nil {
			fmt.Println("FTX close error:", err)
		}

		if err := partnerBinance.Close("Thanks!"); err != nil {
			fmt.Println("Binance close error:", err)
		}

		select {
		case <-time.After(time.Second * 2):
			fmt.Println("No errors after closing")
		case err := <-partnerFTX.ErrorC():
			fmt.Println("FTX Error after closing", err)
		case err := <-partnerBinance.ErrorC():
			fmt.Println("Binance Error after closing", err)
		}
	}

	defer closeAll()

	for {
		select {
		case errSignal := <-partnerFTX.ErrorC():
			log.Printf("FTX Error received %v\n", errSignal)
			return // Or continue?
		case errSignal := <-partnerBinance.ErrorC():
			log.Printf("Binance Error received %v\n", errSignal)
			return // Or continue?
		case <-interrupt:
			log.Println("User interrupt")
			return
		}
	}
}
