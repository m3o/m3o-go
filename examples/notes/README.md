# Notes

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Notes/api](https://m3o.com/Notes/api).

Endpoints:

## Update

Update a note


[https://m3o.com/notes/api#Update](https://m3o.com/notes/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Update a note
func UpdateAnote() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Update(&notes.UpdateRequest{
		Note: &notes.Note{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	Text: "Updated note text",
	Title: "Update Note",
	},

	})
	fmt.Println(rsp, err)
}
```
## Delete

Delete a note


[https://m3o.com/notes/api#Delete](https://m3o.com/notes/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Delete a note
func DeleteAnote() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Delete(&notes.DeleteRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
}
```
## Events

Subscribe to notes events


[https://m3o.com/notes/api#Events](https://m3o.com/notes/api#Events)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Subscribe to notes events
func SubscribeToEvents() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Events(&notes.EventsRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
}
```
## Create

Create a new note


[https://m3o.com/notes/api#Create](https://m3o.com/notes/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Create a new note
func CreateAnote() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Create(&notes.CreateRequest{
		Text: "This is my note",
Title: "New Note",

	})
	fmt.Println(rsp, err)
}
```
## Read

Read a note


[https://m3o.com/notes/api#Read](https://m3o.com/notes/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// Read a note
func ReadAnote() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.Read(&notes.ReadRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
}
```
## List

List all the notes


[https://m3o.com/notes/api#List](https://m3o.com/notes/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// List all the notes
func ListAllNotes() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.List(&notes.ListRequest{
		
	})
	fmt.Println(rsp, err)
}
```
