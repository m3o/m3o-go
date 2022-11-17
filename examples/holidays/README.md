# Holidays

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/holidays/api](https://m3o.com/holidays/api).

Endpoints:

## List

List the holiday dates for a given country and year


[https://m3o.com/holidays/api#List](https://m3o.com/holidays/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/holidays"
)

// List the holiday dates for a given country and year
func GetHolidays() {
	holidaysService := holidays.NewHolidaysService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := holidaysService.List(&holidays.ListRequest{
		Year: 2022,

	})
	fmt.Println(rsp, err)
	
}
```
## Countries

Get the list of countries that are supported by this API


[https://m3o.com/holidays/api#Countries](https://m3o.com/holidays/api#Countries)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/holidays"
)

// Get the list of countries that are supported by this API
func ListCountries() {
	holidaysService := holidays.NewHolidaysService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := holidaysService.Countries(&holidays.CountriesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
