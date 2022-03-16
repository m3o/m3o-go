package main

import (
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// List all events
func main() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.List(&analytics.ListRequest{})
	fmt.Println(rsp, err)
}
