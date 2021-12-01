package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Deploy a group of functions
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Deploy(&function.DeployRequest{
		Entrypoint: "helloworld",
		Name:       "helloworld",
		Repo:       "github.com/m3o/nodejs-function-example",
		Runtime:    "nodejs14",
	})
	fmt.Println(rsp, err)

}
