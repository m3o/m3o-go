package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Create your own NFT (coming soon)
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Create(&nft.CreateRequest{
		Description: "The epic monkey island character",
		Name:        "Guybrush Threepwood",
	})
	fmt.Println(rsp, err)

}
