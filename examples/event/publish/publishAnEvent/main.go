package main

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Publish a event to the event stream.
func main() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
			"id":   "1",
			"type": "signup",
			"user": "john",
		},
		Topic: "user",
	})
	fmt.Println(rsp, err)

}
