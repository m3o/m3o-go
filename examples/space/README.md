# Space

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Space/api](https://m3o.com/Space/api).

Endpoints:

## Delete

Delete an object


[https://m3o.com/space/api#Delete](https://m3o.com/space/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Delete an object
func DeleteAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Delete(&space.DeleteRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List the objects in the space


[https://m3o.com/space/api#List](https://m3o.com/space/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// List the objects in the space
func ListObjectsWithPrefix() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.List(&space.ListRequest{
		Prefix: "images/",

	})
	fmt.Println(rsp, err)
	
}
```
## Head

Retrieve meta information about an object


[https://m3o.com/space/api#Head](https://m3o.com/space/api#Head)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Retrieve meta information about an object
func HeadAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Head(&space.HeadRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read/download the object


[https://m3o.com/space/api#Read](https://m3o.com/space/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Read/download the object
func ReadAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Read(&space.ReadRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create an object. Returns error if object with this name already exists. If you want to update an existing object use the `Update` endpoint
You need to send the request as a multipart/form-data rather than the usual application/json
with each parameter as a form field.


[https://m3o.com/space/api#Create](https://m3o.com/space/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Create an object. Returns error if object with this name already exists. If you want to update an existing object use the `Update` endpoint
// You need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func CreateAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Create(&space.CreateRequest{
		Name: "images/file.jpg",
Object: "<file bytes>",
Visibility: "public",

	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update an object. If an object with this name does not exist, creates a new one.
You need to send the request as a multipart/form-data rather than the usual application/json
with each parameter as a form field.


[https://m3o.com/space/api#Update](https://m3o.com/space/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Update an object. If an object with this name does not exist, creates a new one.
// You need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func UpdateAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Update(&space.UpdateRequest{
		Name: "images/file.jpg",
Object: "<file bytes>",
Visibility: "public",

	})
	fmt.Println(rsp, err)
	
}
```
