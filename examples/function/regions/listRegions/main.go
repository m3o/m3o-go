package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Return a list of supported regions
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Regions(&function.RegionsRequest{})
	fmt.Println(rsp, err)
}
