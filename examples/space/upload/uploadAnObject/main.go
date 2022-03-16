package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/space"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Space.Upload(&space.UploadRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)
}
