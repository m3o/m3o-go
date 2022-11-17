package secret

import (
	"go.m3o.com/client"
)

type Secret interface {
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Get(*GetRequest) (*GetResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Set(*SetRequest) (*SetResponse, error)
}

func NewSecretService(token string) *SecretService {
	return &SecretService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SecretService struct {
	client *client.Client
}

// Delete a secret. If key not found a success response is returned.
func (t *SecretService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("secret", "Delete", request, rsp)

}

// Get a secret by key.
func (t *SecretService) Get(request *GetRequest) (*GetResponse, error) {

	rsp := &GetResponse{}
	return rsp, t.client.Call("secret", "Get", request, rsp)

}

// List all the available secrets
func (t *SecretService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("secret", "List", request, rsp)

}

// Set a secret. Overwrites any existing value already set.
func (t *SecretService) Set(request *SetRequest) (*SetResponse, error) {

	rsp := &SetResponse{}
	return rsp, t.client.Call("secret", "Set", request, rsp)

}

type DeleteRequest struct {
	// The key to delete
	Key string `json:"key,omitempty"`
	// Optional path
	Path string `json:"path,omitempty"`
}

type DeleteResponse struct {
}

type GetRequest struct {
	// The key to retrieve
	Key string `json:"key,omitempty"`
	// Optional path
	Path string `json:"path,omitempty"`
}

type GetResponse struct {
	// time of creation
	Created string `json:"created,omitempty"`
	// The key e.g foo
	Key string `json:"key,omitempty"`
	// Path of value e.g bar/baz
	Path string `json:"path,omitempty"`
	// time of update
	Updated string `json:"updated,omitempty"`
	// The value e.g cat
	Value string `json:"value,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	Keys []string `json:"keys,omitempty"`
}

type SetRequest struct {
	// The key to update
	Key string `json:"key,omitempty"`
	// Optional path e.g bar/baz
	Path string `json:"path,omitempty"`
	// The value to set
	Value string `json:"value,omitempty"`
}

type SetResponse struct {
}
