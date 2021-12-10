package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Upload a large object. Returns a time limited presigned URL to be used for uploading the object
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Upload(&space.UploadRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)
}
