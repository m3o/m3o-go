package main

import (
	"fmt"
	"os"

	"go.m3o.com/bitcoin"
)

// Get the price of bitcoin
func main() {
	bitcoinService := bitcoin.NewBitcoinService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := bitcoinService.Price(&bitcoin.PriceRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}
