package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Update(&db.UpdateRequest{
		Record: map[string]interface{}{
			"age": 43,
			"id":  "1",
		},
		Table: "example",
	})
	fmt.Println(rsp, err)
}
