package main

import (
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Update a comment
func main() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Update(&comments.UpdateRequest{
		Comment: &comments.Comment{
			Id:      "63c0cdf8-2121-11ec-a881-0242e36f037a",
			Subject: "Update Comment",
			Text:    "Updated comment text",
		},
	})
	fmt.Println(rsp, err)
}
