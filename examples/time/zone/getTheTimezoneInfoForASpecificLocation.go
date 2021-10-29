package example

import (
	"fmt"
	"os"

	"go.m3o.com/time"
)

// Get the timezone info for a specific location
func GetTheTimezoneInfoForAspecificLocation() {
	timeService := time.NewTimeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := timeService.Zone(&time.ZoneRequest{
		Location: "London",
	})
	fmt.Println(rsp, err)
}
