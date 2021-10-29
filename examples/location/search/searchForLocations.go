package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/location"
)

// Search for entities in a given radius
func SearchForLocations() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Search(&location.SearchRequest{
		Center: &location.Point{
			Latitude:  51.511061,
			Longitude: -0.120022,
		},
		NumEntities: 10,
		Radius:      100,
		Type:        "bike",
	})
	fmt.Println(rsp, err)
}
