package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stock"
)

// Get the last price for a given stock ticker
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stock.Price(&stock.PriceRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp, err)
}
