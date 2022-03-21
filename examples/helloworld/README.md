# Helloworld

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/helloworld/api](https://m3o.com/helloworld/api).

Endpoints:

## Stream

Stream a personalised Hello message


[https://m3o.com/helloworld/api#Stream](https://m3o.com/helloworld/api#Stream)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Stream a personalised Hello message
func StreamHelloworld() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	
	stream, err := helloworldService.Stream(&helloworld.StreamRequest{
		Messages: 10,
Name: "John",

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
## Call

Return a personalised Hello message


[https://m3o.com/helloworld/api#Call](https://m3o.com/helloworld/api#Call)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/helloworld"
)

// Return a personalised Hello message
func CallHelloworld() {
	helloworldService := helloworld.NewHelloworldService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := helloworldService.Call(&helloworld.CallRequest{
		Name: "John",

	})
	fmt.Println(rsp, err)
	
}
```
