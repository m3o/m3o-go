package main

import (
	"fmt"
	"os"

	"go.m3o.com/translate"
)

// Basic text translation
func main() {
	translateService := translate.NewTranslateService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := translateService.Text(&translate.TextRequest{
		Content: "hello",
		Format:  "text",
		Model:   "nmt",
		Source:  "en",
		Target:  "fr",
	})
	fmt.Println(rsp, err)
}
