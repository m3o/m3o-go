package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/crypto"
)

// Get news related to a currency
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Crypto.News(&crypto.NewsRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)
}
