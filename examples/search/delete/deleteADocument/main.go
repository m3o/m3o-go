package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete a document given its ID
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Delete(&search.DeleteRequest{
		Id:    "1234",
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
