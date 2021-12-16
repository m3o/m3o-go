package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Update a function. Downloads the source, builds and redeploys
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Update(&function.UpdateRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
