package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/bitcoin"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Bitcoin.Transaction(&bitcoin.TransactionRequest{
		Hash: "f854aebae95150b379cc1187d848d58225f3c4157fe992bcd166f58bd5063449",
	})
	fmt.Println(rsp, err)
}
