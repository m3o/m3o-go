package main

import (
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Create a new note
func main() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Create(&notes.CreateRequest{
		Text:  "This is my note",
		Title: "New Note",
	})
	fmt.Println(rsp, err)

}
