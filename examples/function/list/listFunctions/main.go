package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// List all the deployed functions
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.List(&function.ListRequest{})
	fmt.Println(rsp, err)
}
