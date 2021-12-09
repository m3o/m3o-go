package main

import (
	"fmt"
	"os"

	"go.m3o.com/routing"
)

// Turn by turn directions from a start point to an end point including maneuvers and bearings
func main() {
	routingService := routing.NewRoutingService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := routingService.Directions(&routing.DirectionsRequest{
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
