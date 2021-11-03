package main

import (
	"fmt"
	"os"

	"go.m3o.com/quran"
)

// Search the Quran for any form of query or questions
func main() {
	quranService := quran.NewQuranService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := quranService.Search(&quran.SearchRequest{
		Query: "messenger",
	})
	fmt.Println(rsp, err)
}
