package example

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Subscribe to messages for a given topic.
func SubscribeToAtopic() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Subscribe(&event.SubscribeRequest{
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
