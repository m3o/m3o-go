package main

import (
	"fmt"
	"os"

	"go.m3o.com/youtube"
)

// Embed a YouTube video
func main() {
	youtubeService := youtube.NewYoutubeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := youtubeService.Embed(&youtube.EmbedRequest{
		Url: "https://www.youtube.com/watch?v=GWRWZu7XsJ0",
	})
	fmt.Println(rsp, err)
}
