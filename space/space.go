package space

import (
	"go.m3o.com/client"
)

type Space interface {
	Create(*CreateRequest) (*CreateResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Download(*DownloadRequest) (*DownloadResponse, error)
	Head(*HeadRequest) (*HeadResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
	Upload(*UploadRequest) (*UploadResponse, error)
}

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

// Create an object. Returns error if object with this name already exists. Max object size of 10MB, see Upload endpoint for larger objects. If you want to update an existing object use the `Update` endpoint
func (t *SpaceService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("space", "Create", request, rsp)

}

// Delete an object from space
func (t *SpaceService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("space", "Delete", request, rsp)

}

// Download an object via a presigned url
func (t *SpaceService) Download(request *DownloadRequest) (*DownloadResponse, error) {

	rsp := &DownloadResponse{}
	return rsp, t.client.Call("space", "Download", request, rsp)

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

// Read an object in space
func (t *SpaceService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("space", "Read", request, rsp)

}

// Update an object. If an object with this name does not exist, creates a new one.
func (t *SpaceService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("space", "Update", request, rsp)

}

// Upload a large object (> 10MB). Returns a time limited presigned URL to be used for uploading the object
func (t *SpaceService) Upload(request *UploadRequest) (*UploadResponse, error) {

	rsp := &UploadResponse{}
	return rsp, t.client.Call("space", "Upload", request, rsp)

}

type CreateRequest struct {
	// The name of the object. Use forward slash delimiter to implement a nested directory-like structure e.g. images/foo.jpg
	Name string `json:"name,omitempty"`
	// The contents of the object. Either base64 encoded if sending request as application/json or raw bytes if using multipart/form-data format
	Object string `json:"object,omitempty"`
	// Who can see this object? "public" or "private", defaults to "private"
	Visibility string `json:"visibility,omitempty"`
}

type CreateResponse struct {
	// A public URL to access the object if visibility is "public"
	Url string `json:"url,omitempty"`
}

type DeleteRequest struct {
	// Name of the object
	Name string `json:"name,omitempty"`
}

type DeleteResponse struct {
}

type DownloadRequest struct {
	// name of object
	Name string `json:"name,omitempty"`
}

type DownloadResponse struct {
	// presigned url
	Url string `json:"url,omitempty"`
}

type HeadObject struct {
	// when was this created
	Created string `json:"created,omitempty"`
	// when was this last modified
	Modified string `json:"modified,omitempty"`
	Name     string `json:"name,omitempty"`
	// URL to access the object if it is public
	Url string `json:"url,omitempty"`
	// is this public or private
	Visibility string `json:"visibility,omitempty"`
}

type HeadRequest struct {
	// name of the object
	Name string `json:"name,omitempty"`
}

type HeadResponse struct {
	Object *HeadObject `json:"object,omitempty"`
}

type ListObject struct {
	Created string `json:"created,omitempty"`
	// when was this last modified
	Modified   string `json:"modified,omitempty"`
	Name       string `json:"name,omitempty"`
	Url        string `json:"url,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type ListRequest struct {
	// optional prefix for the name e.g. to return all the objects in the images directory pass images/
	Prefix string `json:"prefix,omitempty"`
}

type ListResponse struct {
	Objects []ListObject `json:"objects,omitempty"`
}

type ReadRequest struct {
	// name of the object
	Name string `json:"name,omitempty"`
}

type ReadResponse struct {
	// The object itself
	Object *SpaceObject `json:"object,omitempty"`
}

type SpaceObject struct {
	// when was this created
	Created string `json:"created,omitempty"`
	// the data within the object
	Data string `json:"data,omitempty"`
	// when was this last modified
	Modified string `json:"modified,omitempty"`
	// name of object
	Name string `json:"name,omitempty"`
	// URL to access the object if it is public
	Url string `json:"url,omitempty"`
	// is this public or private
	Visibility string `json:"visibility,omitempty"`
}

type UpdateRequest struct {
	// The name of the object. Use forward slash delimiter to implement a nested directory-like structure e.g. images/foo.jpg
	Name string `json:"name,omitempty"`
	// The contents of the object. Either base64 encoded if sending request as application/json or raw bytes if using multipart/form-data format
	Object string `json:"object,omitempty"`
	// Who can see this object? "public" or "private", defaults to "private"
	Visibility string `json:"visibility,omitempty"`
}

type UpdateResponse struct {
	// A public URL to access the object if visibility is "public"
	Url string `json:"url,omitempty"`
}

type UploadRequest struct {
	Name string `json:"name,omitempty"`
	// is this object public or private
	Visibility string `json:"visibility,omitempty"`
}

type UploadResponse struct {
	// a presigned url to be used for uploading. To use the URL call it with HTTP PUT and pass the object as the request data
	Url string `json:"url,omitempty"`
}
