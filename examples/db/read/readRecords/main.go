package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

// Read data from a table. Lookup can be by ID or via querying any field in the record.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Read(&db.ReadRequest{
		Query: "age == 43",
		Table: "example",
	})
	fmt.Println(rsp, err)
}
