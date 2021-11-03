# Time

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Time/api](https://m3o.com/Time/api).

Endpoints:

## Zone

Get the timezone info for a specific location


[https://m3o.com/time/api#Zone](https://m3o.com/time/api#Zone)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/time"
)

// Get the timezone info for a specific location
func GetTheTimezoneInfoForAspecificLocation() {
	timeService := time.NewTimeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := timeService.Zone(&time.ZoneRequest{
		Location: "London",

	})
	fmt.Println(rsp, err)
}
```
## Now

Get the current time


[https://m3o.com/time/api#Now](https://m3o.com/time/api#Now)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/time"
)

// Get the current time
func ReturnsCurrentTimeOptionallyWithLocation() {
	timeService := time.NewTimeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := timeService.Now(&time.NowRequest{
		
	})
	fmt.Println(rsp, err)
}
```
