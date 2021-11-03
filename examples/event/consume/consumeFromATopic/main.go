package main

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Consume events from a given topic.
func main() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Consume(&event.ConsumeRequest{
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
