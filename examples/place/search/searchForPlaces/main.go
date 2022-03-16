package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/place"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Place.Search(&place.SearchRequest{
		Location: "51.5074577,-0.1297515",
		Query:    "food",
	})
	fmt.Println(rsp, err)
}
