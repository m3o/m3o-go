package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Get the info for a deployed function
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Describe(&function.DescribeRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)

}
