package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/translate"
)

// Basic text translation
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Translate.Text(&translate.TextRequest{
		Content: "hello",
		Format:  "text",
		Model:   "nmt",
		Source:  "en",
		Target:  "fr",
	})
	fmt.Println(rsp, err)
}
