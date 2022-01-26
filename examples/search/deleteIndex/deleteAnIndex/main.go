package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Delete an index by name
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.DeleteIndex(&search.DeleteIndexRequest{
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
