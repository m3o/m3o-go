package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/crypto"
)

// Returns the full list of supported symbols
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Crypto.Symbols(&crypto.SymbolsRequest{})
	fmt.Println(rsp, err)
}
