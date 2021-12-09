package main

import (
	"fmt"
	"os"

	"go.m3o.com/twitter"
)

// Get the timeline for a given user
func main() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.Timeline(&twitter.TimelineRequest{
		Limit:    1,
		Username: "m3oservices",
	})
	fmt.Println(rsp, err)
}
