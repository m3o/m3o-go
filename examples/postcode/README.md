# Postcode

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/postcode/api](https://m3o.com/postcode/api).

Endpoints:

## Validate

Validate a postcode.


[https://m3o.com/postcode/api#Validate](https://m3o.com/postcode/api#Validate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/postcode"
)

// Validate a postcode.
func ReturnArandomPostcodeAndItsInformation() {
	postcodeService := postcode.NewPostcodeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := postcodeService.Validate(&postcode.ValidateRequest{
		Postcode: "SW1A 2AA",

	})
	fmt.Println(rsp, err)
	
}
```
## Lookup

Lookup a postcode to retrieve the related region, county, etc


[https://m3o.com/postcode/api#Lookup](https://m3o.com/postcode/api#Lookup)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/postcode"
)

// Lookup a postcode to retrieve the related region, county, etc
func LookupPostcode() {
	postcodeService := postcode.NewPostcodeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := postcodeService.Lookup(&postcode.LookupRequest{
		Postcode: "SW1A 2AA",

	})
	fmt.Println(rsp, err)
	
}
```
## Random

Return a random postcode and its related info


[https://m3o.com/postcode/api#Random](https://m3o.com/postcode/api#Random)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/postcode"
)

// Return a random postcode and its related info
func ReturnArandomPostcodeAndItsInformation() {
	postcodeService := postcode.NewPostcodeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := postcodeService.Random(&postcode.RandomRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
