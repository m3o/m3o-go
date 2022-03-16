package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cache"
)

// Increment a value (if it's a number). If key not found it is equivalent to set.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cache.Increment(&cache.IncrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp, err)
}
