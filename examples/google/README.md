# Google

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Google/api](https://m3o.com/Google/api).

Endpoints:

## Search

Search for videos on Google


[https://m3o.com/google/api#Search](https://m3o.com/google/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/google"
)

// Search for videos on Google
func SearchForVideos() {
	googleService := google.NewGoogleService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := googleService.Search(&google.SearchRequest{
		Query: "how to make donuts",

	})
	fmt.Println(rsp, err)
}
```
