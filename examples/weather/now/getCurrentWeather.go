package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/weather"
)

// Get the current weather report for a location by postcode, city, zip code, ip address
func GetCurrentWeather() {
	weatherService := weather.NewWeatherService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := weatherService.Now(&weather.NowRequest{
		Location: "london",
	})
	fmt.Println(rsp, err)
}
