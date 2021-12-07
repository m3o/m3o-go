# Movie

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/Movie/api](https://m3o.com/Movie/api).

Endpoints:

## Search

Search for movies by simple text search


[https://m3o.com/movie/api#Search](https://m3o.com/movie/api#Search)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/movie"
)

// Search for movies by simple text search
func SearchForMovies() {
	movieService := movie.NewMovieService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := movieService.Search(&movie.SearchRequest{
		Language: "en-US",
Page: 1,
Query: "inception",
Region: "US",
Year: 2010,

	})
	fmt.Println(rsp, err)
	
}
```
