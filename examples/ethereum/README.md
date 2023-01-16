# Ethereum

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/ethereum/api](https://m3o.com/ethereum/api).

Endpoints:

## Balance

Get the balance of an ethereum wallet


[https://m3o.com/ethereum/api#Balance](https://m3o.com/ethereum/api#Balance)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ethereum"
)

// Get the balance of an ethereum wallet
func GetAethereumAddressBalance() {
	ethereumService := ethereum.NewEthereumService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := ethereumService.Balance(&ethereum.BalanceRequest{
		Address: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",

	})
	fmt.Println(rsp, err)
	
}
```
## Broadcast

Broadcast presigned transaction to ethereum network


[https://m3o.com/ethereum/api#Broadcast](https://m3o.com/ethereum/api#Broadcast)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ethereum"
)

// Broadcast presigned transaction to ethereum network
func BroadcastAtransaction() {
	ethereumService := ethereum.NewEthereumService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := ethereumService.Broadcast(&ethereum.BroadcastRequest{
		Hex: "0xf904808000831cfde080",

	})
	fmt.Println(rsp, err)
	
}
```
## Transaction

Get transaction details by hash


[https://m3o.com/ethereum/api#Transaction](https://m3o.com/ethereum/api#Transaction)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/ethereum"
)

// Get transaction details by hash
func GetAethereumTransaction() {
	ethereumService := ethereum.NewEthereumService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := ethereumService.Transaction(&ethereum.TransactionRequest{
		Hash: "0xbc78ab8a9e9a0bca7d0321a27b2c03addeae08ba81ea98b03cd3dd237eabed44",

	})
	fmt.Println(rsp, err)
	
}
```
