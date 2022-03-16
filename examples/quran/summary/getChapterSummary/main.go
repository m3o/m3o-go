package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/quran"
)

// Get a summary for a given chapter (surah)
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Quran.Summary(&quran.SummaryRequest{
		Chapter: 1,
	})
	fmt.Println(rsp, err)
}
