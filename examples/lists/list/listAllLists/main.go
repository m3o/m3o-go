package main

import (
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// List all the lists
func main() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.List(&lists.ListRequest{})
	fmt.Println(rsp, err)
}
