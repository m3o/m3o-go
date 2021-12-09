package main

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Consume events from a given topic.
func main() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	stream, err := eventService.Consume(&event.ConsumeRequest{
		Topic: "user",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp)
	}
}
