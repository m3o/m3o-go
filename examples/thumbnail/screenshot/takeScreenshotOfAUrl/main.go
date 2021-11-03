package main

import (
	"fmt"
	"os"

	"go.m3o.com/thumbnail"
)

// Create a thumbnail screenshot by passing in a url, height and width
func main() {
	thumbnailService := thumbnail.NewThumbnailService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := thumbnailService.Screenshot(&thumbnail.ScreenshotRequest{
		Height: 600,
		Url:    "https://m3o.com",
		Width:  600,
	})
	fmt.Println(rsp, err)

}
