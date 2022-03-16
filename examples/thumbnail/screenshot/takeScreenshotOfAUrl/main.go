package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/thumbnail"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Thumbnail.Screenshot(&thumbnail.ScreenshotRequest{
		Height: 600,
		Url:    "https://google.com",
		Width:  600,
	})
	fmt.Println(rsp, err)
}
