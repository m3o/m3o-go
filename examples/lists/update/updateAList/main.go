package main

import (
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Update a list
func main() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Update(&lists.UpdateRequest{
		List: &lists.List{
			Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
		},
	})
	fmt.Println(rsp, err)
}
