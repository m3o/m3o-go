package main

import (
	"fmt"
	"os"

	"go.m3o.com"
	"go.m3o.com/youtube"
)

func main() {
	client := m3o.New(os.Getenv("M3O_API_TOKEN"))
	rsp, err := client.Youtube.Embed(&youtube.EmbedRequest{
		Url: "https://www.youtube.com/watch?v=GWRWZu7XsJ0",
	})
	fmt.Println(rsp, err)
}
