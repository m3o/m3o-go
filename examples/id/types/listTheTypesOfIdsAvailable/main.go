package main

import (
	"fmt"
	"os"

	"go.m3o.com/id"
)

// List the types of IDs available. No query params needed.
func main() {
	idService := id.NewIdService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := idService.Types(&id.TypesRequest{})
	fmt.Println(rsp, err)

}
