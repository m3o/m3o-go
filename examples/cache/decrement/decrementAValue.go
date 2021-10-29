package example

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Decrement a value (if it's a number)
func DecrementAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Decrement(&cache.DecrementRequest{
		Key:   "counter",
		Value: 2,
	})
	fmt.Println(rsp, err)
}
