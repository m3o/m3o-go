package example

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Read data from a table. Lookup can be by ID or via querying any field in the record.
func ReadRecords() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Read(&db.ReadRequest{
		Query: "age == 43",
		Table: "users",
	})
	fmt.Println(rsp, err)
}
