package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/nft"
)

// Return a list of assets
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Nft.Assets(&nft.AssetsRequest{
		Limit: 1,
	})
	fmt.Println(rsp, err)
}
