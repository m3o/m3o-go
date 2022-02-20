package main

import (
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Create a new comment
func main() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Create(&comments.CreateRequest{
		Text: "This is my comment",
	})
	fmt.Println(rsp, err)
}
