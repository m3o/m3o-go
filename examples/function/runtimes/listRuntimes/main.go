package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Return a list of supported runtimes
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Runtimes(&function.RuntimesRequest{})
	fmt.Println(rsp, err)
}
