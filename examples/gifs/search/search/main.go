package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/gifs"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Gifs.Search(&gifs.SearchRequest{
		Limit: 2,
		Query: "dogs",
	})
	fmt.Println(rsp, err)
}
