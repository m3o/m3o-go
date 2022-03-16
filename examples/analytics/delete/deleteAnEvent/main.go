package main

import (
	"fmt"
	"os"

	"go.m3o.com/analytics"
)

// Delete an event
func main() {
	analyticsService := analytics.NewAnalyticsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := analyticsService.Delete(&analytics.DeleteRequest{
		Name: "click",
	})
	fmt.Println(rsp, err)
}
