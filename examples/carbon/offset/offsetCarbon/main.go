package main

import (
	"fmt"
	"os"

	"go.m3o.com/carbon"
)

// Purchase 1kg (0.001 tonnes) of carbon offsets in a single request
func main() {
	carbonService := carbon.NewCarbonService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := carbonService.Offset(&carbon.OffsetRequest{})
	fmt.Println(rsp, err)
}
