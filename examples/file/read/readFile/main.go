package main

import (
	"fmt"
	"os"

	"go.m3o.com/file"
)

// Read a file by path
func main() {
	fileService := file.NewFileService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := fileService.Read(&file.ReadRequest{
		Path:    "/document/text-files/file.txt",
		Project: "examples",
	})
	fmt.Println(rsp, err)

}
