package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/weather"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Weather.Forecast(&weather.ForecastRequest{
		Days:     2,
		Location: "London",
	})
	fmt.Println(rsp, err)
}
