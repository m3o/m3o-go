package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/image"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Image.Convert(&image.ConvertRequest{
		Name: "cat.jpeg",
		Url:  "somewebsite.com/cat.png",
	})
	fmt.Println(rsp, err)
}
