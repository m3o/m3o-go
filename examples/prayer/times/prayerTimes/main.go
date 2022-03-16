package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/prayer"
)

// Get the prayer (salah) times for a location on a given date
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Prayer.Times(&prayer.TimesRequest{
		Location: "london",
	})
	fmt.Println(rsp, err)
}
