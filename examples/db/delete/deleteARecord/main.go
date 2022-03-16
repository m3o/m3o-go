package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/db"
)

// Delete a record in the database by id.
func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Db.Delete(&db.DeleteRequest{
		Id:    "1",
		Table: "example",
	})
	fmt.Println(rsp, err)
}
