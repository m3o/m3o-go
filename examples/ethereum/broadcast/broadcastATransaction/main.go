package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ethereum"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ethereum.Broadcast(&ethereum.BroadcastRequest{
		Hex: "0xf904808000831cfde080",
	})
	fmt.Println(rsp, err)
}
