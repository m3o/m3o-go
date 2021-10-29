package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/twitter"
)

// Get the current global trending topics
func GetTheCurrentGlobalTrendingTopics() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.Trends(&twitter.TrendsRequest{})
	fmt.Println(rsp, err)
}
