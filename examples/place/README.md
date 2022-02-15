# Place

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/place/api](https://m3o.com/place/api).

Endpoints:

## Nearby

Search for places nearby, points of interest and geographic locations


[https://m3o.com/place/api#Nearby](https://m3o.com/place/api#Nearby)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Search for places nearby, points of interest and geographic locations
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
## Search




[https://m3o.com/place/api#Search](https://m3o.com/place/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/place"
)

// 
func SearchForPlaces() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Search(&place.SearchRequest{
		Location: "51.5074577,-0.1297515",
Query: "food",

	})
	fmt.Println(rsp, err)
	
}
```
