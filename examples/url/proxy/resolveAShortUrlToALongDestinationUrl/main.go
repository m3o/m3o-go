package main

import (
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Proxy returns the destination URL of a short URL.
func main() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Proxy(&url.ProxyRequest{})
	fmt.Println(rsp, err)
}
