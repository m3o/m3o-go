package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/crypto"
)

// Get the last price for a given crypto ticker
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Crypto.Price(&crypto.PriceRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}
