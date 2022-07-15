# Lists

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/lists/api](https://m3o.com/lists/api).

Endpoints:

## Create

Create a new list


[https://m3o.com/lists/api#Create](https://m3o.com/lists/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Create a new list
func CreateAlist() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Create(&lists.CreateRequest{
		Items: []string{"This is my list"},

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read a list


[https://m3o.com/lists/api#Read](https://m3o.com/lists/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Read a list
func ReadAlist() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Read(&lists.ReadRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List all the lists


[https://m3o.com/lists/api#List](https://m3o.com/lists/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// List all the lists
func ListAllLists() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.List(&lists.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update a list


[https://m3o.com/lists/api#Update](https://m3o.com/lists/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Update a list
func UpdateAlist() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Update(&lists.UpdateRequest{
		List: &lists.List{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	Items: []string{"Updated list text"},
		},

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a list


[https://m3o.com/lists/api#Delete](https://m3o.com/lists/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Delete a list
func DeleteAlist() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := listsService.Delete(&lists.DeleteRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
	
}
```
## Events

Subscribe to lists events


[https://m3o.com/lists/api#Events](https://m3o.com/lists/api#Events)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/lists"
)

// Subscribe to lists events
func SubscribeToEvents() {
	listsService := lists.NewListsService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := listsService.Events(&lists.EventsRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
			rsp, err := stream.Recv()
			if err != nil {
					fmt.Println(err)
					return
			}

			fmt.Println(rsp)
	}
}
```
