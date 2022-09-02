# Wallet

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/wallet/api](https://m3o.com/wallet/api).

Endpoints:

## Read

Get wallet by id


[https://m3o.com/wallet/api#Read](https://m3o.com/wallet/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Get wallet by id
func ReadAwallet() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Read(&wallet.ReadRequest{
		Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",

	})
	fmt.Println(rsp, err)
	
}
```
## Transactions

List the transactions for a wallet


[https://m3o.com/wallet/api#Transactions](https://m3o.com/wallet/api#Transactions)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// List the transactions for a wallet
func ListTransactions() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Transactions(&wallet.TransactionsRequest{
		Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",

	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create a new wallet


[https://m3o.com/wallet/api#Create](https://m3o.com/wallet/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Create a new wallet
func CreateAnewWallet() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Create(&wallet.CreateRequest{
		Description: "No explanation needed",
Name: "Greatness",

	})
	fmt.Println(rsp, err)
	
}
```
## Debit

Debit a wallet


[https://m3o.com/wallet/api#Debit](https://m3o.com/wallet/api#Debit)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Debit a wallet
func DebitWallet() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Debit(&wallet.DebitRequest{
		Amount: 5,
Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",
Reference: "test debit",
Visible: true,

	})
	fmt.Println(rsp, err)
	
}
```
## List

List your wallets


[https://m3o.com/wallet/api#List](https://m3o.com/wallet/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// List your wallets
func ListWallets() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.List(&wallet.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Transfer

Make a transfer from one wallet to another


[https://m3o.com/wallet/api#Transfer](https://m3o.com/wallet/api#Transfer)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Make a transfer from one wallet to another
func TransferMoney() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Transfer(&wallet.TransferRequest{
		Amount: 5,
Reference: "transfer money",
Visible: true,

	})
	fmt.Println(rsp, err)
	
}
```
## Balance

Get the balance of a wallet


[https://m3o.com/wallet/api#Balance](https://m3o.com/wallet/api#Balance)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Get the balance of a wallet
func GetBalance() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Balance(&wallet.BalanceRequest{
		Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a wallet


[https://m3o.com/wallet/api#Delete](https://m3o.com/wallet/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Delete a wallet
func DeleteAwallet() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Delete(&wallet.DeleteRequest{
		Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",

	})
	fmt.Println(rsp, err)
	
}
```
## Credit

Add credit to a wallet


[https://m3o.com/wallet/api#Credit](https://m3o.com/wallet/api#Credit)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/wallet"
)

// Add credit to a wallet
func CreditWallet() {
	walletService := wallet.NewWalletService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := walletService.Credit(&wallet.CreditRequest{
		Amount: 10,
Id: "b6407edd-2e26-45c0-9e2c-343689bbe5f6",
Reference: "test credit",
Visible: true,

	})
	fmt.Println(rsp, err)
	
}
```
