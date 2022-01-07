package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Index a document i.e. insert a document to search for.
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Index(&search.IndexRequest{
		Document: &search.Document{
			Contents: map[string]interface{}{
				"name":     "John Doe",
				"age":      37,
				"starsign": "Leo",
			},
			Id: "1234",
		},
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
