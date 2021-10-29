package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/gifs"
)

// Search for a GIF
func Search() {
	gifsService := gifs.NewGifsService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := gifsService.Search(&gifs.SearchRequest{
		Limit: 2,
		Query: "dogs",
	})
	fmt.Println(rsp, err)
}
