package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/movie"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Movie.Search(&movie.SearchRequest{
		Language: "en-US",
		Page:     1,
		Query:    "inception",
		Region:   "US",
		Year:     2010,
	})
	fmt.Println(rsp, err)
}
