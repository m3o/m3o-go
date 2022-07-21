package url

import (
	"go.m3o.com/client"
)

type Url interface {
	Delete(*DeleteRequest) (*DeleteResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Resolve(*ResolveRequest) (*ResolveResponse, error)
	Shorten(*ShortenRequest) (*ShortenResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewUrlService(token string) *UrlService {
	return &UrlService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type UrlService struct {
	client *client.Client
}

// Delete a URL
func (t *UrlService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("url", "Delete", request, rsp)

}

// List all the shortened URLs
func (t *UrlService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("url", "List", request, rsp)

}

// Resolve returns the destination URL of a short URL.
func (t *UrlService) Resolve(request *ResolveRequest) (*ResolveResponse, error) {

	rsp := &ResolveResponse{}
	return rsp, t.client.Call("url", "Resolve", request, rsp)

}

// Shorten a URL
func (t *UrlService) Shorten(request *ShortenRequest) (*ShortenResponse, error) {

	rsp := &ShortenResponse{}
	return rsp, t.client.Call("url", "Shorten", request, rsp)

}

// Update the destination for a short URL
func (t *UrlService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("url", "Update", request, rsp)

}

type DeleteRequest struct {
	ShortUrl string `json:"shortURL,omitempty"`
}

type DeleteResponse struct {
}

type ListRequest struct {
	// filter by short URL, optional
	ShortUrl string `json:"shortURL,omitempty"`
}

type ListResponse struct {
	UrlPairs *URLPair `json:"urlPairs,omitempty"`
}

type ResolveRequest struct {
	// short url ID, without the domain, eg. if your short URL is
	// `m3o.one/u/someshorturlid` then pass in `someshorturlid`
	ShortUrl string `json:"shortURL,omitempty"`
}

type ResolveResponse struct {
	DestinationUrl string `json:"destinationURL,omitempty"`
}

type ShortenRequest struct {
	// the url to shorten
	DestinationUrl string `json:"destinationURL,omitempty"`
}

type ShortenResponse struct {
	// the shortened url
	ShortUrl string `json:"shortURL,omitempty"`
}

type URLPair struct {
	// time of creation
	Created string `json:"created,omitempty"`
	// destination url
	DestinationUrl string `json:"destinationURL,omitempty"`
	// shortened url
	ShortUrl string `json:"shortURL,omitempty"`
}

type UpdateRequest struct {
	// the destination to update to
	DestinationUrl string `json:"destinationURL,omitempty"`
	// the short url to update
	ShortUrl string `json:"shortURL,omitempty"`
}

type UpdateResponse struct {
}
