package main

import (
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Get the latest quote for the forex
func main() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.Quote(&forex.QuoteRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
