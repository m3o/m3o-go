package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Reserve function names and resources beyond free quota
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Reserve(&function.ReserveRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)
}
