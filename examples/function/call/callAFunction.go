package example

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Call a function by name
func CallAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Call(&function.CallRequest{
		Name:    "my-first-func",
		Request: map[string]interface{}{},
	})
	fmt.Println(rsp, err)
}
