# Chat

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/chat/api](https://m3o.com/chat/api).

Endpoints:

## Leave

Leave a group


[https://m3o.com/chat/api#Leave](https://m3o.com/chat/api#Leave)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Leave a group
func LeaveAgroup() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Leave(&chat.LeaveRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create a new group


[https://m3o.com/chat/api#Create](https://m3o.com/chat/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Create a new group
func CreateAnewChat() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Create(&chat.CreateRequest{
		Description: "The general group",
Name: "general",

	})
	fmt.Println(rsp, err)
	
}
```
## List

List available chats


[https://m3o.com/chat/api#List](https://m3o.com/chat/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// List available chats
func ListGroups() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.List(&chat.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Invite

Invite a user to a group


[https://m3o.com/chat/api#Invite](https://m3o.com/chat/api#Invite)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Invite a user to a group
func InviteAuser() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Invite(&chat.InviteRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## History

List the messages in a chat


[https://m3o.com/chat/api#History](https://m3o.com/chat/api#History)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// List the messages in a chat
func GetChatHistory() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.History(&chat.HistoryRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a group


[https://m3o.com/chat/api#Delete](https://m3o.com/chat/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Delete a group
func DeleteAchat() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Delete(&chat.DeleteRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## Send

Connect to a chat to receive a stream of messages
Send a message to a chat


[https://m3o.com/chat/api#Send](https://m3o.com/chat/api#Send)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Connect to a chat to receive a stream of messages
// Send a message to a chat
func SendAmessage() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Send(&chat.SendRequest{
		Client: "web",
Subject: "Random",
Text: "Hey whats up?",

	})
	fmt.Println(rsp, err)
	
}
```
## Join

Join a group


[https://m3o.com/chat/api#Join](https://m3o.com/chat/api#Join)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Join a group
func JoinAgroup() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := chatService.Join(&chat.JoinRequest{
		
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
## Kick

Kick a user from a group


[https://m3o.com/chat/api#Kick](https://m3o.com/chat/api#Kick)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/chat"
)

// Kick a user from a group
func KickAuserFromAgroup() {
	chatService := chat.NewChatService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := chatService.Kick(&chat.KickRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
