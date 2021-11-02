package example

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Publish a message to the event stream.
func PublishAmessage() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
			"user": "john",
			"id":   "1",
			"type": "signup",
		},
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
