# Joke

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Joke/api](https://m3o.com/Joke/api).

Endpoints:

## Random

Get a random joke


[https://m3o.com/joke/api#Random](https://m3o.com/joke/api#Random)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/joke"
)

// Get a random joke
func GetRandomNjokes() {
	jokeService := joke.NewJokeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := jokeService.Random(&joke.RandomRequest{
		Count: 3,

	})
	fmt.Println(rsp, err)
	
}
```
