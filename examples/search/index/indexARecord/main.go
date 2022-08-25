package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/search"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Search.Index(&search.IndexRequest{
		Data: map[string]interface{}{
			"name":     "John Doe",
			"age":      37,
			"starsign": "Leo",
		},
		Index: "customers",
	})
	fmt.Println(rsp, err)
}
