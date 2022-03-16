package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/currency"
)

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Currency.Convert(&currency.ConvertRequest{
		From: "USD",
		To:   "GBP",
	})
	fmt.Println(rsp, err)
}
