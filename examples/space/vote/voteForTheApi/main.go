package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Vote to have the Space api launched faster!
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Vote(&space.VoteRequest{
		Message: "Launch it!",
	})
	fmt.Println(rsp, err)

}
