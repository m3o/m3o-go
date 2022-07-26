package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/wallet"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Wallet.Debit(&wallet.DebitRequest{
		Amount:    5,
		Id:        "b6407edd-2e26-45c0-9e2c-343689bbe5f6",
		Reference: "test debit",
		Visible:   true,
	})
	fmt.Println(rsp, err)
}
