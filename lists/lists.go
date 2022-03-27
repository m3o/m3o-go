package lists

import (
	"go.m3o.com/client"
)

type Lists interface {
	Create(*CreateRequest) (*CreateResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Events(*EventsRequest) (*EventsResponseStream, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewListsService(token string) *ListsService {
	return &ListsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ListsService struct {
	client *client.Client
}

// Create a new list
func (t *ListsService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("lists", "Create", request, rsp)

}

// Delete a list
func (t *ListsService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("lists", "Delete", request, rsp)

}

// Subscribe to lists events
func (t *ListsService) Events(request *EventsRequest) (*EventsResponseStream, error) {
	stream, err := t.client.Stream("lists", "Events", request)
	if err != nil {
		return nil, err
	}
	return &EventsResponseStream{
		stream: stream,
	}, nil

}

type EventsResponseStream struct {
	stream *client.Stream
}

func (t *EventsResponseStream) Recv() (*EventsResponse, error) {
	var rsp EventsResponse
	if err := t.stream.Recv(&rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// List all the lists
func (t *ListsService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("lists", "List", request, rsp)

}

// Read a list
func (t *ListsService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("lists", "Read", request, rsp)

}

// Update a list
func (t *ListsService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("lists", "Update", request, rsp)

}

type CreateRequest struct {
	// list items
	Items []string `json:"items,omitempty"`
	// list name
	Name string `json:"name,omitempty"`
}

type CreateResponse struct {
	// The created list
	List *List `json:"list,omitempty"`
}

type DeleteRequest struct {
	// specify the id of the list
	Id string `json:"id,omitempty"`
}

type DeleteResponse struct {
	List *List `json:"list,omitempty"`
}

type EventsRequest struct {
	// optionally specify a list id
	Id string `json:"id,omitempty"`
}

type EventsResponse struct {
	// the event which occured; create, delete, update
	Event string `json:"event,omitempty"`
	// the list which the operation occured on
	List *List `json:"list,omitempty"`
}

type List struct {
	// time at which the list was created
	Created string `json:"created,omitempty"`
	// unique id for the list, generated if not specified
	Id string `json:"id,omitempty"`
	// items within the list
	Items []string `json:"items,omitempty"`
	// name of the list
	Name string `json:"name,omitempty"`
	// time at which the list was updated
	Updated string `json:"updated,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	// the list of lists
	Lists []List `json:"lists,omitempty"`
}

type ReadRequest struct {
	// the list id
	Id string `json:"id,omitempty"`
}

type ReadResponse struct {
	// The list
	List *List `json:"list,omitempty"`
}

type UpdateRequest struct {
	List *List `json:"list,omitempty"`
}

type UpdateResponse struct {
	List *List `json:"list,omitempty"`
}
