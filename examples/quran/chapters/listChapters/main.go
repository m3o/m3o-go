package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/quran"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Quran.Chapters(&quran.ChaptersRequest{
		Language: "en",
	})
	fmt.Println(rsp, err)
}
