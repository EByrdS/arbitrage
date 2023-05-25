package main

import (
	"fmt"

	"github.com/arbitrage/apis/ftx"
)

func main() {
	client := ftx.New()
	// response, err := client.GetAllBalances()
	response, err := client.GetDepositAddress("XRP", "")

	// client := binance.New()
	// response, err := client.GetSystemStatus()
	// response, err := client.GetDepositAddress("XRP", "")
	// response, err := client.GetAccountInformation()

	if err != nil {
		fmt.Printf("Error!\n%+v\n", err)
		return
	} else {
		fmt.Printf("Success!\n%+v\n", response)
	}
}
