package main

import (
	"fmt"
	"os"

	"go.m3o.com/crypto"
)

// Returns the full list of supported symbols
func main() {
	cryptoService := crypto.NewCryptoService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cryptoService.Symbols(&crypto.SymbolsRequest{})
	fmt.Println(rsp, err)
}
