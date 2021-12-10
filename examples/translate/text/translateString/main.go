package main

import (
	"fmt"
	"os"

	"go.m3o.com/translate"
)

// TextRequest is the basic edition request
func main() {
	translateService := translate.NewTranslateService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := translateService.Text(&translate.TextRequest{
		Format: "text",
		Model:  "nmt",
		Source: "en",
		Target: "zh",
	})
	fmt.Println(rsp, err)
}
