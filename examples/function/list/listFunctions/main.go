package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// List all the deployed functions
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.List(&function.ListRequest{})
	fmt.Println(rsp, err)
}
