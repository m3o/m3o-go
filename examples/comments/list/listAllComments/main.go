package main

import (
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// List all the comments
func main() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.List(&comments.ListRequest{})
	fmt.Println(rsp, err)
}
