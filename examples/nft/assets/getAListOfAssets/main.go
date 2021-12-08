package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Return a list of NFT assets
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Assets(&nft.AssetsRequest{
		Limit: 1,
	})
	fmt.Println(rsp, err)

}
