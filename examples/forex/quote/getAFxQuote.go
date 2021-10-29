package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/forex"
)

// Get the latest quote for the forex
func GetAfxQuote() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.Quote(&forex.QuoteRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
