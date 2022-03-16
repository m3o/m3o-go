package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/ip"
)

// Lookup the geolocation information for an IP address
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Ip.Lookup(&ip.LookupRequest{
		Ip: "93.148.214.31",
	})
	fmt.Println(rsp, err)
}
