package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/notes"
)

// Update a note
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Notes.Update(&notes.UpdateRequest{
		Note: &notes.Note{
			Id:    "63c0cdf8-2121-11ec-a881-0242e36f037a",
			Text:  "Updated note text",
			Title: "Update Note",
		},
	})
	fmt.Println(rsp, err)
}
