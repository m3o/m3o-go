package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

// Update an object. If an object with this name does not exist, creates a new one.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.Update(&space.UpdateRequest{
		Name:       "images/file.jpg",
		Object:     "<file bytes>",
		Visibility: "public",
	})
	fmt.Println(rsp, err)
}
