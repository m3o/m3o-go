package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Rename a table
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.RenameTable(&db.RenameTableRequest{
		From: "events",
		To:   "events_backup",
	})
	fmt.Println(rsp, err)

}
