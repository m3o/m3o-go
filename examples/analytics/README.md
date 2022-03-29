# Analytics

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/analytics/api](https://m3o.com/analytics/api).

Endpoints:

## List

List all events


[https://m3o.com/analytics/api#List](https://m3o.com/analytics/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// List all events
func ListAllEvents() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.List(&analytics.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete an event


[https://m3o.com/analytics/api#Delete](https://m3o.com/analytics/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Delete an event
func DeleteAnEvent() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Delete(&analytics.DeleteRequest{
		Name: "click",

	})
	fmt.Println(rsp, err)
	
}
```
## Track

Track an event, it will be created if it doesn't exist


[https://m3o.com/analytics/api#Track](https://m3o.com/analytics/api#Track)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Track an event, it will be created if it doesn't exist
func TrackAnEvent() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Track(&analytics.TrackRequest{
		Name: "click",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Get a single event


[https://m3o.com/analytics/api#Read](https://m3o.com/analytics/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Get a single event
func ReadAnEvent() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Read(&analytics.ReadRequest{
		Name: "click",

	})
	fmt.Println(rsp, err)
	
}
```
