package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Search for documents in a given in index
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Search(&search.SearchRequest{
		Index: "customers",
		Query: "name == 'John' AND starsign == 'Leo'",
	})
	fmt.Println(rsp, err)
}
