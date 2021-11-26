# Space

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Space/api](https://m3o.com/Space/api).

Endpoints:

## Vote

Vote to have the Space api launched faster!


[https://m3o.com/space/api#Vote](https://m3o.com/space/api#Vote)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/space"
)

// Vote to have the Space api launched faster!
func VoteForTheApi() {
	spaceService := space.NewSpaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := spaceService.Vote(&space.VoteRequest{
		Message: "Launch it!",

	})
	fmt.Println(rsp, err)
	
}
```
