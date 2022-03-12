# Comments

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/comments/api](https://m3o.com/comments/api).

Endpoints:

## List

List all the comments


[https://m3o.com/comments/api#List](https://m3o.com/comments/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// List all the comments
func ListAllComments() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.List(&comments.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Update

Update a comment


[https://m3o.com/comments/api#Update](https://m3o.com/comments/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Update a comment
func UpdateAcomment() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Update(&comments.UpdateRequest{
		Comment: &comments.Comment{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",
	Subject: "Update Comment",
	Text: "Updated comment text",
	},

	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a comment


[https://m3o.com/comments/api#Delete](https://m3o.com/comments/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Delete a comment
func DeleteAcomment() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Delete(&comments.DeleteRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
	
}
```
## Events

Subscribe to comments events


[https://m3o.com/comments/api#Events](https://m3o.com/comments/api#Events)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Subscribe to comments events
func SubscribeToEvents() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := commentsService.Events(&comments.EventsRequest{
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
## Create

Create a new comment


[https://m3o.com/comments/api#Create](https://m3o.com/comments/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Create a new comment
func CreateAcomment() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Create(&comments.CreateRequest{
		Text: "This is my comment",

	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read a comment


[https://m3o.com/comments/api#Read](https://m3o.com/comments/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/comments"
)

// Read a comment
func ReadAcomment() {
	commentsService := comments.NewCommentsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := commentsService.Read(&comments.ReadRequest{
		Id: "63c0cdf8-2121-11ec-a881-0242e36f037a",

	})
	fmt.Println(rsp, err)
	
}
```
