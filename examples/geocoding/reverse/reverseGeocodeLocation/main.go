package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/geocoding"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Geocoding.Reverse(&geocoding.ReverseRequest{
		Latitude:  51.5123064,
		Longitude: -0.1216235,
	})
	fmt.Println(rsp, err)
}
