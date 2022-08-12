package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ethereum"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ethereum.Balance(&ethereum.BalanceRequest{
		Address: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
	})
	fmt.Println(rsp, err)
}
