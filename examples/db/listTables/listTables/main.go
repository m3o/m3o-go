package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// List tables in the DB
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.ListTables(&db.ListTablesRequest{})
	fmt.Println(rsp, err)
}
