package main

import (
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Track an event, it will be created if it doesn't exist
func main() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Track(&analytics.TrackRequest{
		Name: "click",
	})
	fmt.Println(rsp, err)
}
