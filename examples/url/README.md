# Url

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/url/api](https://m3o.com/url/api).

Endpoints:

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
## Update

Update the destination for a short URL


[https://m3o.com/url/api#Update](https://m3o.com/url/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Update the destination for a short URL
func UpdateAshortUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Update(&url.UpdateRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Shorten

Shorten a URL


[https://m3o.com/url/api#Shorten](https://m3o.com/url/api#Shorten)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Shorten a URL
func ShortenAlongUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Shorten(&url.ShortenRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Resolve

Resolve returns the destination URL of a short URL.


[https://m3o.com/url/api#Resolve](https://m3o.com/url/api#Resolve)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Resolve returns the destination URL of a short URL.
func ResolveAshortUrlToAlongDestinationUrl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Resolve(&url.ResolveRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
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
## Create

Create a URL


[https://m3o.com/url/api#Create](https://m3o.com/url/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/url"
)

// Create a URL
func CreateAurl() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.Create(&url.CreateRequest{
		Id: "my-site",

	})
	fmt.Println(rsp, err)
	
}
```
