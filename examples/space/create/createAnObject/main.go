package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Create an object. Returns error if object with this name already exists. If you want to update an existing object use the `Update` endpoint
// You need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Create(&space.CreateRequest{
		Name:       "images/file.jpg",
		Object:     "<file bytes>",
		Visibility: "public",
	})
	fmt.Println(rsp, err)

}
