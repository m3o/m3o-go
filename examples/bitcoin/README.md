# Bitcoin

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/bitcoin/api](https://m3o.com/bitcoin/api).

Endpoints:

## Price

Get the price of bitcoin


[https://m3o.com/bitcoin/api#Price](https://m3o.com/bitcoin/api#Price)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/bitcoin"
)

// Get the price of bitcoin
func GetBitcoinPrice() {
	bitcoinService := bitcoin.NewBitcoinService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := bitcoinService.Price(&bitcoin.PriceRequest{
		Symbol: "BTCUSD",

	})
	fmt.Println(rsp, err)
	
}
```
## Balance

Get the BTC balance of an address


[https://m3o.com/bitcoin/api#Balance](https://m3o.com/bitcoin/api#Balance)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/bitcoin"
)

// Get the BTC balance of an address
func GetAbitcoinAddressBalance() {
	bitcoinService := bitcoin.NewBitcoinService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := bitcoinService.Balance(&bitcoin.BalanceRequest{
		Address: "1MDUoxL1bGvMxhuoDYx6i11ePytECAk9QK",

	})
	fmt.Println(rsp, err)
	
}
```
