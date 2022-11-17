package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/secret"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Secret.Delete(&secret.DeleteRequest{
		Key: "foo",
	})
	fmt.Println(rsp, err)
}
