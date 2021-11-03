package main

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Increment a value (if it's a number). If key not found it is equivalent to set.
func main() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Increment(&cache.IncrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp, err)

}
