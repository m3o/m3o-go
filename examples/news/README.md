# News

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/News/api](https://m3o.com/News/api).

Endpoints:

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
