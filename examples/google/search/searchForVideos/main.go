package main

import (
	"fmt"
	"os"

	"go.m3o.com/google"
)

// Search for videos on Google
func main() {
	googleService := google.NewGoogleService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := googleService.Search(&google.SearchRequest{
		Query: "how to make donuts",
	})
	fmt.Println(rsp, err)
}
