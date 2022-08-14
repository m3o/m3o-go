# Place

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/place/api](https://m3o.com/place/api).

Endpoints:

## Search

Search for places by text query


[https://m3o.com/place/api#Search](https://m3o.com/place/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Search for places by text query
func SearchForPlaces() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Search(&place.SearchRequest{
		Location: "51.5074577,-0.1297515",
Query: "food",

	})
	fmt.Println(rsp, err)
	
}
```
## Nearby

Find places nearby using a location


[https://m3o.com/place/api#Nearby](https://m3o.com/place/api#Nearby)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Find places nearby using a location
func FindPlacesNearby() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Nearby(&place.NearbyRequest{
		Keyword: "tesco",
Location: "51.5074577,-0.1297515",
Type: "store",

	})
	fmt.Println(rsp, err)
	
}
```
