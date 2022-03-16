package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/twitter"
)

// Get a user's twitter profile
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Twitter.User(&twitter.UserRequest{
		Username: "crufter",
	})
	fmt.Println(rsp, err)
}
