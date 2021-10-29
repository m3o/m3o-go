package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/stream"
)

// Subscribe to messages for a given topic.
func SubscribeToAtopic() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.Subscribe(&stream.SubscribeRequest{
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
