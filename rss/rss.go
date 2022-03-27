package rss

import (
	"go.m3o.com/client"
)

type Rss interface {
	Add(*AddRequest) (*AddResponse, error)
	Feed(*FeedRequest) (*FeedResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Remove(*RemoveRequest) (*RemoveResponse, error)
}

func NewRssService(token string) *RssService {
	return &RssService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type RssService struct {
	client *client.Client
}

// Add a new RSS feed with a name, url, and category
func (t *RssService) Add(request *AddRequest) (*AddResponse, error) {

	rsp := &AddResponse{}
	return rsp, t.client.Call("rss", "Add", request, rsp)

}

// Get an RSS feed by name. If no name is given, all feeds are returned. Default limit is 25 entries.
func (t *RssService) Feed(request *FeedRequest) (*FeedResponse, error) {

	rsp := &FeedResponse{}
	return rsp, t.client.Call("rss", "Feed", request, rsp)

}

// List the saved RSS fields
func (t *RssService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("rss", "List", request, rsp)

}

// Remove an RSS feed by name
func (t *RssService) Remove(request *RemoveRequest) (*RemoveResponse, error) {

	rsp := &RemoveResponse{}
	return rsp, t.client.Call("rss", "Remove", request, rsp)

}

type AddRequest struct {
	// category to add e.g news
	Category string `json:"category,omitempty"`
	// rss feed name
	// eg. a16z
	Name string `json:"name,omitempty"`
	// rss feed url
	// eg. http://a16z.com/feed/
	Url string `json:"url,omitempty"`
}

type AddResponse struct {
}

type Entry struct {
	// article content
	Content string `json:"content,omitempty"`
	// data of the entry
	Date string `json:"date,omitempty"`
	// the rss feed where it came from
	Feed string `json:"feed,omitempty"`
	// unique id of the entry
	Id string `json:"id,omitempty"`
	// rss feed url of the entry
	Link string `json:"link,omitempty"`
	// article summary
	Summary string `json:"summary,omitempty"`
	// title of the entry
	Title string `json:"title,omitempty"`
}

type Feed struct {
	// category of the feed e.g news
	Category string `json:"category,omitempty"`
	// unique id
	Id string `json:"id,omitempty"`
	// rss feed name
	// eg. a16z
	Name string `json:"name,omitempty"`
	// rss feed url
	// eg. http://a16z.com/feed/
	Url string `json:"url,omitempty"`
}

type FeedRequest struct {
	// limit entries returned
	Limit int64 `json:"limit,string,omitempty"`
	// rss feed name
	Name string `json:"name,omitempty"`
	// offset entries
	Offset int64 `json:"offset,string,omitempty"`
}

type FeedResponse struct {
	Entries []Entry `json:"entries,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	Feeds []Feed `json:"feeds,omitempty"`
}

type RemoveRequest struct {
	// rss feed name
	// eg. a16z
	Name string `json:"name,omitempty"`
}

type RemoveResponse struct {
}
