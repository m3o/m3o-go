package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/routing"
)

// Get the eta for a route from origin to destination. The eta is an estimated time based on car routes
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Routing.Eta(&routing.EtaRequest{
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
