package main

import (
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// List all the active channels
func main() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.ListChannels(&stream.ListChannelsRequest{})
	fmt.Println(rsp, err)

}
