package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/rss"
)

// Remove an RSS feed by name
func RemoveAfeed() {
	rssService := rss.NewRssService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := rssService.Remove(&rss.RemoveRequest{
		Name: "bbc",
	})
	fmt.Println(rsp, err)
}
