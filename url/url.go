package url

import (
	"go.m3o.com/client"
)

type Url interface {
	Create(*CreateRequest) (*CreateResponse, error)
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

// Create a URL
func (t *UrlService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("url", "Create", request, rsp)

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

type CreateRequest struct {
	// destination url
	DestinationUrl string `json:"destinationURL,omitempty"`
	// a unique id e.g uuid or my-url
	Id string `json:"id,omitempty"`
}

type CreateResponse struct {
	Url *URLPair `json:"url,omitempty"`
}

type DeleteRequest struct {
	// delete by id
	Id string `json:"id,omitempty"`
	// delete by shortURL
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
	// short url to resolve
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
	// The number of times the short URL has been resolved
	HitCount int64 `json:"hitCount,string,omitempty"`
	// url id
	Id string `json:"id,omitempty"`
	// shortened url
	ShortUrl string `json:"shortURL,omitempty"`
}

type UpdateRequest struct {
	// the destination to update to
	DestinationUrl string `json:"destinationURL,omitempty"`
	// update by id
	Id string `json:"id,omitempty"`
	// update by short url
	ShortUrl string `json:"shortURL,omitempty"`
}

type UpdateResponse struct {
}
