package main

import (
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping an IP address
func main() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Ip(&ping.IpRequest{
		Address: "google.com",
	})
	fmt.Println(rsp, err)
}
