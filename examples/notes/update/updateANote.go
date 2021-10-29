package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/notes"
)

// Update a note
func UpdateAnote() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Update(&notes.UpdateRequest{
		Note: &notes.Note{
			Id:    "63c0cdf8-2121-11ec-a881-0242e36f037a",
			Text:  "Updated note text",
			Title: "Update Note",
		},
	})
	fmt.Println(rsp, err)
}
