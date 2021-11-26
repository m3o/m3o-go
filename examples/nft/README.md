# Nft

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Nft/api](https://m3o.com/Nft/api).

Endpoints:

## Vote

Vote to have the NFT api launched faster!


[https://m3o.com/nft/api#Vote](https://m3o.com/nft/api#Vote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Vote to have the NFT api launched faster!
func VoteForTheApi() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Vote(&nft.VoteRequest{
		Message: "Launch it!",

	})
	fmt.Println(rsp, err)
	
}
```
