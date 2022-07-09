<p align="center">
	<a href="https://discord.gg/TBR9bRjd6Z">
		<img src="https://discordapp.com/api/guilds/861917584437805127/widget.png?style=banner2" alt="Discord Banner"/>
	</a>
</p>

---

# M3O Go Client [![godoc](https://godoc.org/github.com/m3o/m3o-go?status.svg)](https://godoc.org/github.com/m3o/m3o-go) [![Go Report Card](https://goreportcard.com/badge/github.com/m3o/m3o-go)](https://goreportcard.com/report/github.com/m3o/m3o-go) [![Apache 2.0 License](https://img.shields.io/github/license/m3o/m3o-go)](https://github.com/m3o/m3o-go/blob/master/LICENSE)

This is the Go client to access APIs on the M3O Platform

## Usage

Call a service using the generated client. Populate the `M3O_API_TOKEN` environment variable. 

Import the package and initialise the service with your API token.

```go
package main

import(
        "fmt"
        "os"

	"go.m3o.com"
        hw "go.m3o.com/helloworld"
)

func main() {
	token := os.Getenv("M3O_API_TOKEN")
	client := m3o.New(token)

	rsp, err := client.Helloworld.Call(&hw.CallRequest{
                Name: "Alice",

        })
	
        fmt.Println(rsp, err)
}
```

## Generic Client

The generic client enables you to call any endpoint by name with freeform request/response types.

```go
package main

import (
    "fmt"
    "os"

    "go.m3o.com/client"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Msg string `json:"msg"`
}

var (
	token, _ = os.Getenv("TOKEN")
)

func main() {
	c := client.NewClient(nil)

	// set your api token
	c.SetToken(token)

   	req := &Request{
		Name: "John",
	}
	
	var rsp Response

	if err := c.Call("helloworld", "call", req, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(rsp)
}
```

## Streaming

The client supports streaming but is not yet code generated. Use the following for streaming endpoints.
 
```go
package main

import (
	"fmt"

	"go.m3o.com/client"
)

type Request struct {
	Count string `json:"count"`
}

type Response struct {
	Count string `json:"count"`
}

var (
	token, _ = os.Getenv("TOKEN")
)

func main() {
	c := client.NewClient(nil)

	// set your api token
	c.SetToken(token)
	
	stream, err := c.Stream("streams", "subscribe", Request{Count: "10"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var rsp Response
		if err := stream.Recv(&rsp); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("got", rsp.Count)
	}
}
```
