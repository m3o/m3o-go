package main

import (
	"fmt"
	"os"

	"go.m3o.com/stock"
)

// Get the last quote for the stock
func main() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.Quote(&stock.QuoteRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp, err)
}
