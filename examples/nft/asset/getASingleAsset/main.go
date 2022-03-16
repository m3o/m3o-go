package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/nft"
)

// Get a single asset by the contract
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Nft.Asset(&nft.AssetRequest{})
	fmt.Println(rsp, err)
}
