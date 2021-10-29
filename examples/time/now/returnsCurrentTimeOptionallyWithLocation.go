package example

import (
	"fmt"
	"os"

	"go.m3o.com/time"
)

// Get the current time
func ReturnsCurrentTimeOptionallyWithLocation() {
	timeService := time.NewTimeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := timeService.Now(&time.NowRequest{})
	fmt.Println(rsp, err)
}
