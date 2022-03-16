package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/comments"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Comments.Create(&comments.CreateRequest{
		Text: "This is my comment",
	})
	fmt.Println(rsp, err)
}
