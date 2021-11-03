package main

import (
	"fmt"
	"os"

	"go.m3o.com/quran"
)

// List the Chapters (surahs) of the Quran
func main() {
	quranService := quran.NewQuranService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := quranService.Chapters(&quran.ChaptersRequest{
		Language: "en",
	})
	fmt.Println(rsp, err)

}
