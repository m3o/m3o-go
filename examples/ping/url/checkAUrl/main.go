package main

import (
	"fmt"
	"os"

	"go.m3o.com/ping"
)

// Ping a HTTP URL
func main() {
	pingService := ping.NewPingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := pingService.Url(&ping.UrlRequest{
		Address: "google.com",
	})
	fmt.Println(rsp, err)
}
