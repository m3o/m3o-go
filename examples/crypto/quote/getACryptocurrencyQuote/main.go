package main

import (
	"fmt"
	"os"

	"go.m3o.com/crypto"
)

// Get the last quote for a given crypto ticker
func main() {
	cryptoService := crypto.NewCryptoService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cryptoService.Quote(&crypto.QuoteRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)

}
