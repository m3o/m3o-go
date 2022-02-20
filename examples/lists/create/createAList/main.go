package main

import (
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Create a new list
func main() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Create(&lists.CreateRequest{})
	fmt.Println(rsp, err)
}
