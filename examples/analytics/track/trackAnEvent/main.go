package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/analytics"
)

// Track an event, it will be created if it doesn't exist
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Analytics.Track(&analytics.TrackRequest{
		Name: "click",
	})
	fmt.Println(rsp, err)
}
