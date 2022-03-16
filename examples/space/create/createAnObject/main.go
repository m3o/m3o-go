package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

// Create an object. Returns error if object with this name already exists. Max object size of 10MB, see Upload endpoint for larger objects. If you want to update an existing object use the `Update` endpoint
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.Create(&space.CreateRequest{
		Name:       "images/file.jpg",
		Object:     "<file bytes>",
		Visibility: "public",
	})
	fmt.Println(rsp, err)
}
