package example

import (
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Specify the note to events
func SubscribeToEvents() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Events(&notes.EventsRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	})
	fmt.Println(rsp, err)
}
