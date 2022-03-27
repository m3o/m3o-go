package google

import (
	"go.m3o.com/client"
)

type Google interface {
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewGoogleService(token string) *GoogleService {
	return &GoogleService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type GoogleService struct {
	client *client.Client
}

// Search for videos on Google
func (t *GoogleService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("google", "Search", request, rsp)

}

type SearchRequest struct {
	// Query to search for
	Query string `json:"query,omitempty"`
}

type SearchResponse struct {
	// List of results for the query
	Results []SearchResult `json:"results,omitempty"`
}

type SearchResult struct {
	// abridged version of this search result’s URL, e.g. www.exampe.com
	DisplayUrl string `json:"display_url,omitempty"`
	// id of the result
	Id string `json:"id,omitempty"`
	// kind of result; "search"
	Kind string `json:"kind,omitempty"`
	// the result snippet
	Snippet string `json:"snippet,omitempty"`
	// title of the result
	Title string `json:"title,omitempty"`
	// the full url for the result
	Url string `json:"url,omitempty"`
}
