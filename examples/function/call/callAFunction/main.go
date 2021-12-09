package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Call a function by name
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Call(&function.CallRequest{
		Name:    "helloworld",
		Request: map[string]interface{}{},
	})
	fmt.Println(rsp, err)
}
