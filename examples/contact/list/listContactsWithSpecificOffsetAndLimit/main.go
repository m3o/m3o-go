package main

import (
	"fmt"
	"os"

	"go.m3o.com/contact"
)

//
func main() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.List(&contact.ListRequest{
		Limit:  1,
		Offset: 1,
	})
	fmt.Println(rsp, err)
}
