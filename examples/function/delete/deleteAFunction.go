package example

import (
	"fmt"
	"os"

	"go.m3o.com/function"
)

// Delete a function by name
func DeleteAfunction() {
	functionService := function.NewFunctionService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := functionService.Delete(&function.DeleteRequest{
		Name:    "my-first-func",
	})
	fmt.Println(rsp, err)
}
