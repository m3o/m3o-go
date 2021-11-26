package main

import (
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Vote to have the NFT api launched faster!
func main() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Vote(&nft.VoteRequest{
		Message: "Launch it!",
	})
	fmt.Println(rsp, err)

}
