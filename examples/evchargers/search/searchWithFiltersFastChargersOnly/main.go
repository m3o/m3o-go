package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/evchargers"
)

// Search by giving a coordinate and a max distance, or bounding box and optional filters
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Evchargers.Search(&evchargers.SearchRequest{
		Distance: 2000,
		Levels:   []string{"3"},
		Location: &evchargers.Coordinates{
			Latitude:  51.53336351319885,
			Longitude: -0.0252,
		},
	})
	fmt.Println(rsp, err)
}
