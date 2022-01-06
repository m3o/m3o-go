# Nft

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Nft/api](https://m3o.com/Nft/api).

Endpoints:

## Create

Create your own NFT (coming soon)


[https://m3o.com/nft/api#Create](https://m3o.com/nft/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Create your own NFT (coming soon)
func CreateAnNft() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Create(&nft.CreateRequest{
		Description: "The epic monkey island character",
Name: "Guybrush Threepwood",

	})
	fmt.Println(rsp, err)
	
}
```
## Collections

Get a list of collections


[https://m3o.com/nft/api#Collections](https://m3o.com/nft/api#Collections)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Get a list of collections
func ListCollections() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Collections(&nft.CollectionsRequest{
		Limit: 1,

	})
	fmt.Println(rsp, err)
	
}
```
## Assets

Return a list of assets


[https://m3o.com/nft/api#Assets](https://m3o.com/nft/api#Assets)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/nft"
)

// Return a list of assets
func GetAlistOfAssets() {
	nftService := nft.NewNftService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := nftService.Assets(&nft.AssetsRequest{
		Limit: 1,

	})
	fmt.Println(rsp, err)
	
}
```
