package main

import (
	"fmt"
	"os"

	"go.m3o.com/postcode"
)

// Validate a postcode.
func main() {
	postcodeService := postcode.NewPostcodeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := postcodeService.Validate(&postcode.ValidateRequest{
		Postcode: "SW1A 2AA",
	})
	fmt.Println(rsp, err)

}
