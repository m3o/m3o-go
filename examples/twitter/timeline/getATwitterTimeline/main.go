package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/twitter"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Twitter.Timeline(&twitter.TimelineRequest{
		Limit:    1,
		Username: "m3oservices",
	})
	fmt.Println(rsp, err)
}
