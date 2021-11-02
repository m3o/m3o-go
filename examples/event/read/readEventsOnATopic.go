package example

import (
	"fmt"
	"os"

	"go.m3o.com/event"
)

// Read stored events
func ReadEventsOnAtopic() {
	eventService := event.NewEventService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := eventService.Read(&event.ReadRequest{
		Topic: "user",
	})
	fmt.Println(rsp, err)
}
