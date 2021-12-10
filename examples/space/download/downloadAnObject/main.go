package main

import (
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Download an object via a presigned url
func main() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Download(&space.DownloadRequest{
		Name: "images/file.jpg",
	})
	fmt.Println(rsp, err)
}
