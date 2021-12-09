package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Delete an object from space
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Delete(&space.DeleteRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)
}
