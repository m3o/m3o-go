package example

import (
	"fmt"
	"os"

	"go.m3o.com/file"
)

// Save a file
func SaveFile() {
	fileService := file.NewFileService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := fileService.Save(&file.SaveRequest{
		File: &file.Record{
			Content: "file content example",
			Path:    "/document/text-files/file.txt",
			Project: "examples",
		},
	})
	fmt.Println(rsp, err)
}
