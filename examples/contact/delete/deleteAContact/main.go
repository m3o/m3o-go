package main

import (
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Delete a contact
func main() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Delete(&contact.DeleteRequest{
		Id: "42e48a3c-6221-11ec-96d2-acde48001122",
	})
	fmt.Println(rsp, err)
}
