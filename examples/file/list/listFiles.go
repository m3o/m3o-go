package example

import (
	"fmt"
	"os"

	"go.m3o.com/file"
)

// List files by their project and optionally a path.
func ListFiles() {
	fileService := file.NewFileService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := fileService.List(&file.ListRequest{
		Project: "examples",
	})
	fmt.Println(rsp, err)
}
