package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/price"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Price.Report(&price.ReportRequest{
		Comment: "Price is not up to date",
		Name:    "bitcoin",
		Symbol:  "BTC",
	})
	fmt.Println(rsp, err)
}
