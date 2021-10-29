# Location

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Location/api](https://m3o.com/Location/api).

Endpoints:

## Save

Save an entity's current position


[https://m3o.com/location/api#Save](https://m3o.com/location/api#Save)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/location"
)

// Save an entity's current position
func SaveAnEntity() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Save(&location.SaveRequest{
		Entity: &location.Entity{
	Id: "1",
	Location: &location.Point{
		Latitude: 51.511061,
		Longitude: -0.120022,
		Timestamp: 1622802761,
},
	Type: "bike",
},

	})
	fmt.Println(rsp, err)
}
```
## Read

Read an entity by its ID


[https://m3o.com/location/api#Read](https://m3o.com/location/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/location"
)

// Read an entity by its ID
func GetLocationById() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Read(&location.ReadRequest{
		Id: "1",

	})
	fmt.Println(rsp, err)
}
```
## Search

Search for entities in a given radius


[https://m3o.com/location/api#Search](https://m3o.com/location/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/location"
)

// Search for entities in a given radius
func SearchForLocations() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Search(&location.SearchRequest{
		Center: &location.Point{
	Latitude: 51.511061,
	Longitude: -0.120022,
	},
NumEntities: 10,
Radius: 100,
Type: "bike",

	})
	fmt.Println(rsp, err)
}
```
