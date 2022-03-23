package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/price"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Price.Add(&price.AddRequest{
		Currency: "USD",
		Name:     "bitcoin",
		Price:    39037.97,
	})
	fmt.Println(rsp, err)
}
