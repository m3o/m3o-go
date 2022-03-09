package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a collection by its slug
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Collection(&nft.CollectionRequest{
		Slug: "doodles-official",
	})
	fmt.Println(rsp, err)
}
