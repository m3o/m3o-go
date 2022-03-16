package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/time"
)

// Get the timezone info for a specific location
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Time.Zone(&time.ZoneRequest{
		Location: "London",
	})
	fmt.Println(rsp, err)
}
