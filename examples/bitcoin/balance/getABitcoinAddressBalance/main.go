package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/bitcoin"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Bitcoin.Balance(&bitcoin.BalanceRequest{
		Address: "1MDUoxL1bGvMxhuoDYx6i11ePytECAk9QK",
	})
	fmt.Println(rsp, err)
}
