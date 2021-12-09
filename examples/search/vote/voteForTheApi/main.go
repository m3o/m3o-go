package main

import (
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Vote to have the Search api launched faster!
func main() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Vote(&search.VoteRequest{
		Message: "Launch it!",
	})
	fmt.Println(rsp, err)
}
