package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// List the objects in space
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.List(&space.ListRequest{
		Prefix: "images/",
	})
	fmt.Println(rsp, err)

}
