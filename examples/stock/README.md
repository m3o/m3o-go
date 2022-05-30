# Stock

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/stock/api](https://m3o.com/stock/api).

Endpoints:

## Quote

Get the last quote for the stock


[https://m3o.com/stock/api#Quote](https://m3o.com/stock/api#Quote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stock"
)

// Get the last quote for the stock
func GetAstockQuote() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.Quote(&stock.QuoteRequest{
		Symbol: "AAPL",

	})
	fmt.Println(rsp, err)
	
}
```
## History

Get the historic open-close for a given day


[https://m3o.com/stock/api#History](https://m3o.com/stock/api#History)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stock"
)

// Get the historic open-close for a given day
func GetHistoricData() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.History(&stock.HistoryRequest{
		Date: "2020-10-01",
Stock: "AAPL",

	})
	fmt.Println(rsp, err)
	
}
```
## Price

Get the last price for a given stock ticker


[https://m3o.com/stock/api#Price](https://m3o.com/stock/api#Price)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/stock"
)

// Get the last price for a given stock ticker
func GetAstockPrice() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.Price(&stock.PriceRequest{
		Symbol: "AAPL",

	})
	fmt.Println(rsp, err)
	
}
```
