package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/rss"
)

// List the saved RSS fields
func ListRssFeeds() {
	rssService := rss.NewRssService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := rssService.List(&rss.ListRequest{})
	fmt.Println(rsp, err)
}
