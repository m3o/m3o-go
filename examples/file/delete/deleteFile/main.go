package main

import (
	"fmt"
	"os"

	"go.m3o.com/file"
)

// Delete a file by project name/path
func main() {
	fileService := file.NewFileService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := fileService.Delete(&file.DeleteRequest{
		Path:    "/document/text-files/file.txt",
		Project: "examples",
	})
	fmt.Println(rsp, err)
}
