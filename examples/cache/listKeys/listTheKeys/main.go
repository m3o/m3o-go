package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cache"
)

// List all the available keys
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cache.ListKeys(&cache.ListKeysRequest{})
	fmt.Println(rsp, err)
}
