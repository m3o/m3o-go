package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ping"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ping.Tcp(&ping.TcpRequest{
		Address: "google.com:80",
	})
	fmt.Println(rsp, err)
}
