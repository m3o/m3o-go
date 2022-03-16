package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/image"
)

// Upload an image by either sending a base64 encoded image to this endpoint or a URL.
// To resize an image before uploading, see the Resize endpoint.
// To use the file parameter you need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Image.Upload(&image.UploadRequest{
		Name: "cat.jpeg",
		Url:  "somewebsite.com/cat.png",
	})
	fmt.Println(rsp, err)
}
