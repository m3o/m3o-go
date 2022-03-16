package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/memegen"
)

// Generate a meme using a template
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Memegen.Generate(&memegen.GenerateRequest{
		Id: "444501",
	})
	fmt.Println(rsp, err)
}
