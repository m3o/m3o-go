# Search

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Search/api](https://m3o.com/Search/api).

Endpoints:

## Vote

Vote to have the Search api launched faster!


[https://m3o.com/search/api#Vote](https://m3o.com/search/api#Vote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/search"
)

// Vote to have the Search api launched faster!
func VoteForTheApi() {
	searchService := search.NewSearchService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := searchService.Vote(&search.VoteRequest{
		Message: "Launch it!",

	})
	fmt.Println(rsp, err)
	
}
```
