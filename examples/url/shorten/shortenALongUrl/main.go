package main

import (
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Shorten a long URL
func main() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Shorten(&url.ShortenRequest{})
	fmt.Println(rsp, err)
}
