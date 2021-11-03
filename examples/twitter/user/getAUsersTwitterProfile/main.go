package main

import (
	"fmt"
	"os"

	"go.m3o.com/twitter"
)

// Get a user's twitter profile
func main() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.User(&twitter.UserRequest{
		Username: "crufter",
	})
	fmt.Println(rsp, err)

}
