package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/file"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.File.List(&file.ListRequest{
		Project: "examples",
	})
	fmt.Println(rsp, err)
}
