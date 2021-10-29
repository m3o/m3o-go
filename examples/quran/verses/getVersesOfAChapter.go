package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/quran"
)

// Lookup the verses (ayahs) for a chapter including
// translation, interpretation and breakdown by individual
// words.
func GetVersesOfAchapter() {
	quranService := quran.NewQuranService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := quranService.Verses(&quran.VersesRequest{
		Chapter: 1,
	})
	fmt.Println(rsp, err)
}
