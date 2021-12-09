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
		From: "examples2",
		To:   "examples3",
	})
	fmt.Println(rsp, err)
}
