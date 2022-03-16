package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/currency"
)

// Codes returns the supported currency codes for the API
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Currency.Codes(&currency.CodesRequest{})
	fmt.Println(rsp, err)
}
