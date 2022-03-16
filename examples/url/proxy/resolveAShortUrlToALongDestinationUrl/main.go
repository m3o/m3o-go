package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/url"
)

// Proxy returns the destination URL of a short URL.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Url.Proxy(&url.ProxyRequest{})
	fmt.Println(rsp, err)
}
