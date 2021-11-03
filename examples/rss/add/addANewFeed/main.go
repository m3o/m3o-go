package main

import (
	"fmt"
	"os"

	"go.m3o.com/rss"
)

// Add a new RSS feed with a name, url, and category
func main() {
	rssService := rss.NewRssService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := rssService.Add(&rss.AddRequest{
		Category: "news",
		Name:     "bbc",
		Url:      "http://feeds.bbci.co.uk/news/rss.xml",
	})
	fmt.Println(rsp, err)
}
