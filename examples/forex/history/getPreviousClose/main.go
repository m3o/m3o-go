package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/forex"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Forex.History(&forex.HistoryRequest{
		Symbol: "GBPUSD",
	})
	fmt.Println(rsp, err)
}
