package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Return the backend url for proxying
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Proxy(&function.ProxyRequest{
		Id: "helloworld",
	})
	fmt.Println(rsp, err)
}
