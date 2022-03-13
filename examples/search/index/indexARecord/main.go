package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Index a record i.e. insert a document to search for.
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Index(&search.IndexRequest{
		Data: map[string]interface{}{
			"name":     "John Doe",
			"age":      37,
			"starsign": "Leo",
		},
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
