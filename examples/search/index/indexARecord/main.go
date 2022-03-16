package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/search"
)

// Index a record i.e. insert a document to search for.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Search.Index(&search.IndexRequest{
		Data: map[string]interface{}{
			"starsign": "Leo",
			"name":     "John Doe",
			"age":      37,
		},
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
