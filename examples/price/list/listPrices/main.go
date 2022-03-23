package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/price"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Price.List(&price.ListRequest{
		Currency: "GBP",
	})
	fmt.Println(rsp, err)
}
