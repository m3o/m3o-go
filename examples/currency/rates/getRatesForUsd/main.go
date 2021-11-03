package main

import (
	"fmt"
	"os"

	"go.m3o.com/currency"
)

// Rates returns the currency rates for a given code e.g USD
func main() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.Rates(&currency.RatesRequest{
		Code: "USD",
	})
	fmt.Println(rsp, err)

}
