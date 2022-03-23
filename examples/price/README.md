# Price

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/price/api](https://m3o.com/price/api).

Endpoints:

## Add

Add a price


[https://m3o.com/price/api#Add](https://m3o.com/price/api#Add)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/price"
)

// Add a price
func AddAprice() {
	priceService := price.NewPriceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := priceService.Add(&price.AddRequest{
		Currency: "USD",
Name: "bitcoin",
Price: 39037.97,

	})
	fmt.Println(rsp, err)
	
}
```
## Get

Get the price of anything


[https://m3o.com/price/api#Get](https://m3o.com/price/api#Get)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/price"
)

// Get the price of anything
func GetThePrice() {
	priceService := price.NewPriceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := priceService.Get(&price.GetRequest{
		Currency: "USD",
Name: "bitcoin",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List prices for a given currency


[https://m3o.com/price/api#List](https://m3o.com/price/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/price"
)

// List prices for a given currency
func ListPrices() {
	priceService := price.NewPriceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := priceService.List(&price.ListRequest{
		Currency: "GBP",

	})
	fmt.Println(rsp, err)
	
}
```
## Index

Get the index for available prices


[https://m3o.com/price/api#Index](https://m3o.com/price/api#Index)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/price"
)

// Get the index for available prices
func GetTheIndex() {
	priceService := price.NewPriceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := priceService.Index(&price.IndexRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
