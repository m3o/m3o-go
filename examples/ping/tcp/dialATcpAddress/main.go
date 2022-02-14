package main

import (
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping a TCP port is open
func main() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Tcp(&ping.TcpRequest{
		Address: "google.com:80",
	})
	fmt.Println(rsp, err)
}
