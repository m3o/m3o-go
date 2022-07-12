# Url

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/url/api](https://m3o.com/url/api).

Endpoints:

## Delete

Delete a URL


[https://m3o.com/url/api#Delete](https://m3o.com/url/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Delete a URL
func DeleteAshortenedUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Delete(&url.DeleteRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## List

List all the shortened URLs


[https://m3o.com/url/api#List](https://m3o.com/url/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// List all the shortened URLs
func ListYourShortenedUrls() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.List(&url.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Shorten

Shorten a long URL


[https://m3o.com/url/api#Shorten](https://m3o.com/url/api#Shorten)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Shorten a long URL
func ShortenAlongUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Shorten(&url.ShortenRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Proxy

Proxy returns the destination URL of a short URL.


[https://m3o.com/url/api#Proxy](https://m3o.com/url/api#Proxy)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Proxy returns the destination URL of a short URL.
func ResolveAshortUrlToAlongDestinationUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Proxy(&url.ProxyRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
