package example

import (
	"fmt"
	"os"

	"github.com/go.m3o.com/url"
)

// List information on all the shortened URLs that you have created
func ListYourShortenedUrls() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.List(&url.ListRequest{})
	fmt.Println(rsp, err)
}
