# Space

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/space/api](https://m3o.com/space/api).

Endpoints:

## List

List the objects in space


[https://m3o.com/space/api#List](https://m3o.com/space/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// List the objects in space
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

Read an object in space


[https://m3o.com/space/api#Read](https://m3o.com/space/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Read an object in space
func ReadAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Read(&space.ReadRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## Download

Download an object via a presigned url


[https://m3o.com/space/api#Download](https://m3o.com/space/api#Download)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Download an object via a presigned url
func DownloadAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Download(&space.DownloadRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## Upload

Upload a large object (> 10MB). Returns a time limited presigned URL to be used for uploading the object


[https://m3o.com/space/api#Upload](https://m3o.com/space/api#Upload)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Upload a large object (> 10MB). Returns a time limited presigned URL to be used for uploading the object
func UploadAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Upload(&space.UploadRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create an object. Returns error if object with this name already exists. Max object size of 10MB, see Upload endpoint for larger objects. If you want to update an existing object use the `Update` endpoint


[https://m3o.com/space/api#Create](https://m3o.com/space/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Create an object. Returns error if object with this name already exists. Max object size of 10MB, see Upload endpoint for larger objects. If you want to update an existing object use the `Update` endpoint
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


[https://m3o.com/space/api#Update](https://m3o.com/space/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Update an object. If an object with this name does not exist, creates a new one.
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
## Delete

Delete an object from space


[https://m3o.com/space/api#Delete](https://m3o.com/space/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Delete an object from space
func DeleteAnObject() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Delete(&space.DeleteRequest{
		Name: "images/file.jpg",

	})
	fmt.Println(rsp, err)
	
}
```
