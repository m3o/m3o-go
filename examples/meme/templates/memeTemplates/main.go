package main

import (
	"fmt"
	"os"

	"go.m3o.com/meme"
)

// List the available templates
func main() {
	memeService := meme.NewMemeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memeService.Templates(&meme.TemplatesRequest{})
	fmt.Println(rsp, err)
}
