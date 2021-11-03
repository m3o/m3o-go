package main

import (
	"fmt"
	"os"

	"go.m3o.com/crypto"
)

// Get news related to a currency
func main() {
	cryptoService := crypto.NewCryptoService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cryptoService.News(&crypto.NewsRequest{
		Symbol: "BTCUSD",
	})
	fmt.Println(rsp, err)

}
