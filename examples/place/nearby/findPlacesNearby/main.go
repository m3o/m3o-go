package main

import (
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Find places nearby using a location
func main() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Nearby(&place.NearbyRequest{
		Keyword:  "tesco",
		Location: "51.5074577,-0.1297515",
		Type:     "store",
	})
	fmt.Println(rsp, err)
}
