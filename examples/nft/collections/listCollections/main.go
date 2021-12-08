package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a list of collections
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Collections(&nft.CollectionsRequest{
		Limit: 1,
	})
	fmt.Println(rsp, err)

}
