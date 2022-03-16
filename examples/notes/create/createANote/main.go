package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/notes"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Notes.Create(&notes.CreateRequest{
		Text:  "This is my note",
		Title: "New Note",
	})
	fmt.Println(rsp, err)
}
