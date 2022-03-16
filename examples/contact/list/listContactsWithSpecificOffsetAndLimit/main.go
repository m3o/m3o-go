package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/contact"
)

// List contacts
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Contact.List(&contact.ListRequest{
		Limit:  1,
		Offset: 1,
	})
	fmt.Println(rsp, err)
}
