package main

import (
	"fmt"
	"os"

	"go.m3o.com/news"
)

// Get the latest news headlines
func main() {
	newsService := news.NewNewsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := newsService.Headlines(&news.HeadlinesRequest{
		Date:     "2021-11-24",
		Language: "en",
		Locale:   "us",
	})
	fmt.Println(rsp, err)

}
