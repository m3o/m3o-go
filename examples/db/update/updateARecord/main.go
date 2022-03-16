package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

// Update a record in the database. Include an "id" in the record to update.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Update(&db.UpdateRequest{
		Record: map[string]interface{}{
			"id":  "1",
			"age": 43,
		},
		Table: "example",
	})
	fmt.Println(rsp, err)
}
