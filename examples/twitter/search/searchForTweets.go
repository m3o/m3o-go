package example

import (
	"fmt"
	"os"

	"go.m3o.com/twitter"
)

// Search for tweets with a simple query
func SearchForTweets() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.Search(&twitter.SearchRequest{
		Query: "cats",
	})
	fmt.Println(rsp, err)
}
