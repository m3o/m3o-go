package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/youtube"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Youtube.Search(&youtube.SearchRequest{
		Query: "donuts",
	})
	fmt.Println(rsp, err)
}
