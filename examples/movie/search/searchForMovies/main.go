package main

import (
	"fmt"
	"os"

	"go.m3o.com/movie"
)

//
func main() {
	movieService := movie.NewMovieService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := movieService.Search(&movie.SearchRequest{
		Language: "en-US",
		Page:     1,
		Query:    "inception",
		Region:   "US",
		Year:     2010,
	})
	fmt.Println(rsp, err)

}
