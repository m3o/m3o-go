# Geocoding

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/geocoding/api](https://m3o.com/geocoding/api).

Endpoints:

## Lookup

Lookup returns a geocoded address including normalized address and gps coordinates. All fields are optional, provide more to get more accurate results


[https://m3o.com/geocoding/api#Lookup](https://m3o.com/geocoding/api#Lookup)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/geocoding"
)

// Lookup returns a geocoded address including normalized address and gps coordinates. All fields are optional, provide more to get more accurate results
func GeocodeAnAddress() {
	geocodingService := geocoding.NewGeocodingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := geocodingService.Lookup(&geocoding.LookupRequest{
		Address: "10 russell st",
City: "london",
Country: "uk",
Postcode: "wc2b",

	})
	fmt.Println(rsp, err)
	
}
```
## Reverse

Reverse lookup an address from gps coordinates


[https://m3o.com/geocoding/api#Reverse](https://m3o.com/geocoding/api#Reverse)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/geocoding"
)

// Reverse lookup an address from gps coordinates
func ReverseGeocodeLocation() {
	geocodingService := geocoding.NewGeocodingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := geocodingService.Reverse(&geocoding.ReverseRequest{
		Latitude: 51.5123064,
Longitude: -0.1216235,

	})
	fmt.Println(rsp, err)
	
}
```
