package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/twitter"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Twitter.Trends(&twitter.TrendsRequest{})
	fmt.Println(rsp, err)
}
