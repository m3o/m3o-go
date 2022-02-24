package main

import (
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// List contacts
func main() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.List(&contact.ListRequest{})
	fmt.Println(rsp, err)
}
