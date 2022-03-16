package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/forex"
)

// Get the latest price for a given forex ticker
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Forex.Price(&forex.PriceRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
