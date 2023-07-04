# News

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/news/api](https://m3o.com/news/api).

Endpoints:

## TopStories

Get the top stories


[https://m3o.com/news/api#TopStories](https://m3o.com/news/api#TopStories)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/news"
)

// Get the top stories
func GetTheTopStories() {
	newsService := news.NewNewsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := newsService.TopStories(&news.TopStoriesRequest{
		Date: "2021-11-24",
Language: "en",
Locale: "us",

	})
	fmt.Println(rsp, err)
	
}
```
## Headlines

Get the latest news headlines


[https://m3o.com/news/api#Headlines](https://m3o.com/news/api#Headlines)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/news"
)

// Get the latest news headlines
func GetNewsHeadlines() {
	newsService := news.NewNewsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := newsService.Headlines(&news.HeadlinesRequest{
		Date: "2021-11-24",
Language: "en",
Locale: "us",

	})
	fmt.Println(rsp, err)
	
}
```
