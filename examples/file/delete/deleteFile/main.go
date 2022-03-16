package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/file"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.File.Delete(&file.DeleteRequest{
		Path:    "/document/text-files/file.txt",
		Project: "examples",
	})
	fmt.Println(rsp, err)
}
