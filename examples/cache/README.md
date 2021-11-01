# Cache

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Cache/api](https://m3o.com/Cache/api).

Endpoints:

## Get

Get an item from the cache by key


[https://m3o.com/cache/api#Get](https://m3o.com/cache/api#Get)

```go
package example

import(
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
```
## Delete

Delete a value from the cache


[https://m3o.com/cache/api#Delete](https://m3o.com/cache/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Delete a value from the cache
func DeleteAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Delete(&cache.DeleteRequest{
		Key: "foo",

	})
	fmt.Println(rsp, err)
}
```
## Increment

Increment a value (if it's a number)


[https://m3o.com/cache/api#Increment](https://m3o.com/cache/api#Increment)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Increment a value (if it's a number)
func IncrementAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Increment(&cache.IncrementRequest{
		Key: "counter",
Value: 2,

	})
	fmt.Println(rsp, err)
}
```
## Decrement

Decrement a value (if it's a number)


[https://m3o.com/cache/api#Decrement](https://m3o.com/cache/api#Decrement)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Decrement a value (if it's a number)
func DecrementAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Decrement(&cache.DecrementRequest{
		Key: "counter",
Value: 2,

	})
	fmt.Println(rsp, err)
}
```
## Set

Set an item in the cache. Overwrites any existing value already set.


[https://m3o.com/cache/api#Set](https://m3o.com/cache/api#Set)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Set an item in the cache. Overwrites any existing value already set.
func SetAvalue() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Set(&cache.SetRequest{
		Key: "foo",
Value: "bar",

	})
	fmt.Println(rsp, err)
}
```
