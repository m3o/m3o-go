package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Retrieve meta information about an object
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Head(&space.HeadRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)

}
