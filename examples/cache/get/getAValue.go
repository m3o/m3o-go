package example

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Get an item from the cache by key
func GetAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Get(&cache.GetRequest{
		Key: "foo",
	})
	fmt.Println(rsp, err)
}
