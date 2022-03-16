package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/quran"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Quran.Search(&quran.SearchRequest{
		Query: "messenger",
	})
	fmt.Println(rsp, err)
}
