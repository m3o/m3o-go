package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/sunnah"
)

// Get a list of available collections. A collection is
// a compilation of hadiths collected and written by an author.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Sunnah.Collections(&sunnah.CollectionsRequest{})
	fmt.Println(rsp, err)
}
