package main

import (
	"fmt"
	"os"

	"go.m3o.com/stream"
)

// Publish a message to the stream. Specify a topic to group messages for a specific topic.
func main() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.Publish(&stream.PublishRequest{
		Message: map[string]interface{}{
			"id":   "1",
			"type": "signup",
			"user": "john",
		},
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
