package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a single asset by the contract
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Asset(&nft.AssetRequest{})
	fmt.Println(rsp, err)
}
