package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/carbon"
)

// Purchase 1KG (0.001 tonne) of carbon offsets in a single request
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Carbon.Offset(&carbon.OffsetRequest{})
	fmt.Println(rsp, err)
}
