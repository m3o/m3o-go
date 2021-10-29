package example

import (
	"fmt"
	"os"

	"go.m3o.com/prayer"
)

// Get the prayer (salah) times for a location on a given date
func PrayerTimes() {
	prayerService := prayer.NewPrayerService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := prayerService.Times(&prayer.TimesRequest{
		Location: "london",
	})
	fmt.Println(rsp, err)
}
