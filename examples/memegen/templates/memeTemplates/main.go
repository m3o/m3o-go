package main

import (
	"fmt"
	"os"

	"go.m3o.com/memegen"
)

// List the available templates
func main() {
	memegenService := memegen.NewMemegenService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := memegenService.Templates(&memegen.TemplatesRequest{})
	fmt.Println(rsp, err)
}
