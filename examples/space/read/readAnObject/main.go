package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Read an object in storage
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Read(&space.ReadRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)

}
