package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/quran"
)

// Get a summary for a given chapter (surah)
func GetChapterSummary() {
	quranService := quran.NewQuranService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := quranService.Summary(&quran.SummaryRequest{
		Chapter: 1,
	})
	fmt.Println(rsp, err)
}
