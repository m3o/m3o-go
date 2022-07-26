package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/wallet"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Wallet.Transfer(&wallet.TransferRequest{
		Amount:    5,
		Reference: "transfer money",
		Visible:   true,
	})
	fmt.Println(rsp, err)
}
