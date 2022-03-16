package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/rss"
)

// Add a new RSS feed with a name, url, and category
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Rss.Add(&rss.AddRequest{
		Category: "news",
		Name:     "bbc",
		Url:      "http://feeds.bbci.co.uk/news/rss.xml",
	})
	fmt.Println(rsp, err)
}
