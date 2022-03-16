package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/memegen"
)

// List the available templates
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Memegen.Templates(&memegen.TemplatesRequest{})
	fmt.Println(rsp, err)
}
