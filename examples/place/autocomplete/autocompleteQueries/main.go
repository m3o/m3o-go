package main

import (
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Autocomplete queries (coming soon)
func main() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Autocomplete(&place.AutocompleteRequest{})
	fmt.Println(rsp, err)
}
