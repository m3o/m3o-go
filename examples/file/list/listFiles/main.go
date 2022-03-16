package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/file"
)

// List files by their project and optionally a path.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.File.List(&file.ListRequest{
		Project: "examples",
	})
	fmt.Println(rsp, err)
}
