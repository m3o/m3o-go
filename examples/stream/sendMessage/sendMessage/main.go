package main

import (
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// Send a message to the stream.
func main() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.SendMessage(&stream.SendMessageRequest{
		Channel: "general",
		Text:    "Hey checkout this tweet https://twitter.com/m3oservices/status/1455291054295498752",
	})
	fmt.Println(rsp, err)
}
