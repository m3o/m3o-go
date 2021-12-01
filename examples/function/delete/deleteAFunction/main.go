package main

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Delete a function by name
func main() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Delete(&function.DeleteRequest{
		Name: "helloworld",
	})
	fmt.Println(rsp, err)

}
