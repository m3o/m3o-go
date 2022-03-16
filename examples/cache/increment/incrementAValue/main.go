package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/cache"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Cache.Increment(&cache.IncrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp, err)
}
