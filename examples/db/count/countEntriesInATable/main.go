package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Count records in a table
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Count(&db.CountRequest{
		Table: "users",
	})
	fmt.Println(rsp, err)

}
