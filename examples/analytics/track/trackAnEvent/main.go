package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/analytics"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Analytics.Track(&analytics.TrackRequest{
		Name: "click",
	})
	fmt.Println(rsp, err)
}
