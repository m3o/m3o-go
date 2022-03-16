package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/routing"
)

// Retrieve a route as a simple list of gps points along with total distance and estimated duration
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Routing.Route(&routing.RouteRequest{
		Destination: &routing.Point{
			Latitude:  52.529407,
			Longitude: 13.397634,
		},
		Origin: &routing.Point{
			Latitude:  52.517037,
			Longitude: 13.38886,
		},
	})
	fmt.Println(rsp, err)
}
