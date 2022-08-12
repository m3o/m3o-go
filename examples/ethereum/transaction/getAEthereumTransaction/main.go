package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ethereum"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ethereum.Transaction(&ethereum.TransactionRequest{
		Hash: "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44",
	})
	fmt.Println(rsp, err)
}
