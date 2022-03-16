package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/event"
)

// Publish a event to the event stream.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Event.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
			"type": "signup",
			"user": "john",
			"id":   "1",
		},
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
