package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/google"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Google.Search(&google.SearchRequest{
		Query: "how to make donuts",
	})
	fmt.Println(rsp, err)
}
