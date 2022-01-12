# Youtube

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/youtube/api](https://m3o.com/youtube/api).

Endpoints:

## Embed

Embed a YouTube video


[https://m3o.com/youtube/api#Embed](https://m3o.com/youtube/api#Embed)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/youtube"
)

// Embed a YouTube video
func EmbedAyoutubeVideo() {
	youtubeService := youtube.NewYoutubeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := youtubeService.Embed(&youtube.EmbedRequest{
		Url: "https://www.youtube.com/watch?v=GWRWZu7XsJ0",

	})
	fmt.Println(rsp, err)
	
}
```
## Search

Search for videos on YouTube


[https://m3o.com/youtube/api#Search](https://m3o.com/youtube/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/youtube"
)

// Search for videos on YouTube
func SearchForVideos() {
	youtubeService := youtube.NewYoutubeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := youtubeService.Search(&youtube.SearchRequest{
		Query: "donuts",

	})
	fmt.Println(rsp, err)
	
}
```
