package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Update an object. If an object with this name does not exist, creates a new one.
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Update(&space.UpdateRequest{
		Name:       "images/file.jpg",
		Object:     "<file bytes>",
		Visibility: "public",
	})
	fmt.Println(rsp, err)
}
