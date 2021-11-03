package main

import (
	"fmt"
	"os"

	"go.m3o.com/quran"
)

// Lookup the verses (ayahs) for a chapter including
// translation, interpretation and breakdown by individual
// words.
func main() {
	quranService := quran.NewQuranService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := quranService.Verses(&quran.VersesRequest{
		Chapter: 1,
	})
	fmt.Println(rsp, err)

}
