# Db

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/db/api](https://m3o.com/db/api).

Endpoints:

## Create

Create a record in the database. Optionally include an "id" field otherwise it's set automatically.


[https://m3o.com/db/api#Create](https://m3o.com/db/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Create a record in the database. Optionally include an "id" field otherwise it's set automatically.
func CreateArecord() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Create(&db.CreateRequest{
		Record: map[string]interface{}{
	"name": "Jane",
	"age": 42,
	"isActive": true,
	"id": "1",
},
Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update a record in the database. Include an "id" in the record to update.


[https://m3o.com/db/api#Update](https://m3o.com/db/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Update a record in the database. Include an "id" in the record to update.
func UpdateArecord() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Update(&db.UpdateRequest{
		Record: map[string]interface{}{
	"age": 43,
	"id": "1",
},
Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## ListTables

List tables in the DB


[https://m3o.com/db/api#ListTables](https://m3o.com/db/api#ListTables)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// List tables in the DB
func ListTables() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.ListTables(&db.ListTablesRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## RenameTable

Rename a table


[https://m3o.com/db/api#RenameTable](https://m3o.com/db/api#RenameTable)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Rename a table
func RenameTable() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.RenameTable(&db.RenameTableRequest{
		From: "examples2",
To: "examples3",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read data from a table. Lookup can be by ID or via querying any field in the record.


[https://m3o.com/db/api#Read](https://m3o.com/db/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Read data from a table. Lookup can be by ID or via querying any field in the record.
func ReadRecords() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Read(&db.ReadRequest{
		Query: "age == 43",
Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a record in the database by id.


[https://m3o.com/db/api#Delete](https://m3o.com/db/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Delete a record in the database by id.
func DeleteArecord() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Delete(&db.DeleteRequest{
		Id: "1",
Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## Truncate

Truncate the records in a table


[https://m3o.com/db/api#Truncate](https://m3o.com/db/api#Truncate)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Truncate the records in a table
func TruncateTable() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Truncate(&db.TruncateRequest{
		Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## DropTable

Drop a table in the DB


[https://m3o.com/db/api#DropTable](https://m3o.com/db/api#DropTable)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Drop a table in the DB
func DropTable() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.DropTable(&db.DropTableRequest{
		Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
## Count

Count records in a table


[https://m3o.com/db/api#Count](https://m3o.com/db/api#Count)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Count records in a table
func CountEntriesInAtable() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Count(&db.CountRequest{
		Table: "example",

	})
	fmt.Println(rsp, err)
	
}
```
