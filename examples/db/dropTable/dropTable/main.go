package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Drop a table in the DB
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.DropTable(&db.DropTableRequest{
		Table: "users",
	})
	fmt.Println(rsp, err)

}
