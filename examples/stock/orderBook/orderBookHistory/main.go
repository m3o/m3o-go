package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/stock"
)

// Get the historic order book and each trade by timestamp
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Stock.OrderBook(&stock.OrderBookRequest{
		Date:  "2020-10-01",
		End:   "2020-10-01T11:00:00Z",
		Limit: 3,
		Start: "2020-10-01T10:00:00Z",
		Stock: "AAPL",
	})
	fmt.Println(rsp, err)
}
