# Forex

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Forex/api](https://m3o.com/Forex/api).

Endpoints:

## Price

Get the latest price for a given forex ticker


[https://m3o.com/forex/api#Price](https://m3o.com/forex/api#Price)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Get the latest price for a given forex ticker
func GetAnFxPrice() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.Price(&forex.PriceRequest{
		Symbol: "GBPUSD",

	})
	fmt.Println(rsp, err)
	
}
```
## Quote

Get the latest quote for the forex


[https://m3o.com/forex/api#Quote](https://m3o.com/forex/api#Quote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Get the latest quote for the forex
func GetAfxQuote() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.Quote(&forex.QuoteRequest{
		Symbol: "GBPUSD",

	})
	fmt.Println(rsp, err)
	
}
```
## History

Returns the data for the previous close


[https://m3o.com/forex/api#History](https://m3o.com/forex/api#History)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/forex"
)

// Returns the data for the previous close
func GetPreviousClose() {
	forexService := forex.NewForexService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := forexService.History(&forex.HistoryRequest{
		Symbol: "GBPUSD",

	})
	fmt.Println(rsp, err)
	
}
```
