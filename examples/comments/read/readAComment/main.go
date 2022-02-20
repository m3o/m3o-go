package main

import (
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Read a comment
func main() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Read(&comments.ReadRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	})
	fmt.Println(rsp, err)
}
