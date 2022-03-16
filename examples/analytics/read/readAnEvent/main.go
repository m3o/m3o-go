package main

import (
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Get a single event
func main() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Read(&analytics.ReadRequest{})
	fmt.Println(rsp, err)
}
