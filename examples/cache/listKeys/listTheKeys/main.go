package main

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// List all the available keys
func main() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.ListKeys(&cache.ListKeysRequest{})
	fmt.Println(rsp, err)
}
