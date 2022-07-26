package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/wallet"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Wallet.Credit(&wallet.CreditRequest{
		Amount:    10,
		Id:        "b6407edd-2e26-45c0-9e2c-343689bbe5f6",
		Reference: "test credit",
		Visible:   true,
	})
	fmt.Println(rsp, err)
}
