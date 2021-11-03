package main

import (
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Get the latest price for a given forex ticker
func main() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.Price(&forex.PriceRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
