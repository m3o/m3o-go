package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cache"
)

// Get an item from the cache by key. If key is not found, an empty response is returned.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cache.Get(&cache.GetRequest{
		Key: "foo",
	})
	fmt.Println(rsp, err)
}
