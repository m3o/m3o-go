package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/sunnah"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Sunnah.Chapters(&sunnah.ChaptersRequest{
		Book:       1,
		Collection: "bukhari",
	})
	fmt.Println(rsp, err)
}
