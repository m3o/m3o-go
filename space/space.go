package space

import (
	"go.m3o.com/client"
)

func NewSpaceService(token string) *SpaceService {
	return &SpaceService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SpaceService struct {
	client *client.Client
}

// Create an object. Returns error if object with this name already exists. If you want to update an existing object use the `Update` endpoint
// You need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func (t *SpaceService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("space", "Create", request, rsp)

}

// Delete an object from space
func (t *SpaceService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("space", "Delete", request, rsp)

}

// Retrieve meta information about an object
func (t *SpaceService) Head(request *HeadRequest) (*HeadResponse, error) {

	rsp := &HeadResponse{}
	return rsp, t.client.Call("space", "Head", request, rsp)

}

// List the objects in space
func (t *SpaceService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("space", "List", request, rsp)

}

// Read an object in space. Use for private objects.
func (t *SpaceService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("space", "Read", request, rsp)

}

// Update an object. If an object with this name does not exist, creates a new one.
// You need to send the request as a multipart/form-data rather than the usual application/json
// with each parameter as a form field.
func (t *SpaceService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("space", "Update", request, rsp)

}

type CreateRequest struct {
	// The name of the object. Use forward slash delimiter to implement a nested directory-like structure e.g. images/foo.jpg
	Name string `json:"name"`
	// The contents of the object
	Object string `json:"object"`
	// Who can see this object? "public" or "private", defaults to "private"
	Visibility string `json:"visibility"`
}

type CreateResponse struct {
	// A public URL to access the object if visibility is "public"
	Url string `json:"url"`
}

type DeleteRequest struct {
	// The name of the object. Use forward slash delimiter to implement a nested directory-like structure e.g. images/foo.jpg
	Name string `json:"name"`
}

type DeleteResponse struct {
}

type HeadObject struct {
	// when was this created
	Created int64 `json:"created,string"`
	// when was this last modified
	Modified int64  `json:"modified,string"`
	Name     string `json:"name"`
	// URL to access the object if it is public
	Url string `json:"url"`
	// is this public or private
	Visibility string `json:"visibility"`
}

type HeadRequest struct {
	// name of the object
	Name string `json:"name"`
}

type HeadResponse struct {
	Object *HeadObject `json:"object"`
}

type ListObject struct {
	// when was this last modified
	Modified int64  `json:"modified,string"`
	Name     string `json:"name"`
	Url      string `json:"url"`
}

type ListRequest struct {
	// optional prefix for the name e.g. to return all the objects in the images directory pass images/
	Prefix string `json:"prefix"`
}

type ListResponse struct {
	Objects []ListObject `json:"objects"`
}

type ReadRequest struct {
	// name of the object
	Name string `json:"name"`
}

type ReadResponse struct {
	// Returns the object as raw data
	Object string `json:"object"`
}

type UpdateRequest struct {
	// The name of the object. Use forward slash delimiter to implement a nested directory-like structure e.g. images/foo.jpg
	Name string `json:"name"`
	// The contents of the object
	Object string `json:"object"`
	// Who can see this object? "public" or "private", defaults to "private"
	Visibility string `json:"visibility"`
}

type UpdateResponse struct {
	// A public URL to access the object if visibility is "public"
	Url string `json:"url"`
}
