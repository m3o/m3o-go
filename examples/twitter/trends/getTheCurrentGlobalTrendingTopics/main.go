package main

import (
	"fmt"
	"os"

	"go.m3o.com/twitter"
)

// Get the current global trending topics
func main() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.Trends(&twitter.TrendsRequest{})
	fmt.Println(rsp, err)

}
