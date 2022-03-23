package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/price"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Price.Get(&price.GetRequest{
		Currency: "USD",
		Name:     "bitcoin",
	})
	fmt.Println(rsp, err)
}
