package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/function"
)

// Return a list of supported runtimes
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Function.Runtimes(&function.RuntimesRequest{})
	fmt.Println(rsp, err)
}
