package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Create an index by name
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.CreateIndex(&search.CreateIndexRequest{
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
