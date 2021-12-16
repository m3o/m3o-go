package example

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Get the info for a deployed function
func DescribeFunctionStatus() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Describe(&function.DescribeRequest{
		Name:    "my-first-func",
	})
	fmt.Println(rsp, err)
}
