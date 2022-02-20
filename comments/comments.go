package comments

import (
	"go.m3o.com/client"
)

type Comments interface {
	Create(*CreateRequest) (*CreateResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Events(*EventsRequest) (*EventsResponseStream, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewCommentsService(token string) *CommentsService {
	return &CommentsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CommentsService struct {
	client *client.Client
}

// Create a new comment
func (t *CommentsService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("comments", "Create", request, rsp)

}

// Delete a comment
func (t *CommentsService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("comments", "Delete", request, rsp)

}

// Subscribe to comments events
func (t *CommentsService) Events(request *EventsRequest) (*EventsResponseStream, error) {
	stream, err := t.client.Stream("comments", "Events", request)
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

// List all the comments
func (t *CommentsService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("comments", "List", request, rsp)

}

// Read a comment
func (t *CommentsService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("comments", "Read", request, rsp)

}

// Update a comment
func (t *CommentsService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("comments", "Update", request, rsp)

}

type Comment struct {
	// time at which the comment was created
	Created string `json:"created"`
	// unique id for the comment, generated if not specified
	Id string `json:"id"`
	// subject of the comment
	Subject string `json:"subject"`
	// text of the comment
	Text string `json:"text"`
	// time at which the comment was updated
	Updated string `json:"updated"`
}

type CreateRequest struct {
	// comment subject
	Subject string `json:"subject"`
	// comment items
	Text string `json:"text"`
}

type CreateResponse struct {
	// The created comment
	Comment *Comment `json:"comment"`
}

type DeleteRequest struct {
	// specify the id of the comment
	Id string `json:"id"`
}

type DeleteResponse struct {
	Comment *Comment `json:"comment"`
}

type EventsRequest struct {
	// optionally specify a comment id
	Id string `json:"id"`
}

type EventsResponse struct {
	// the comment which the operation occured on
	Comment *Comment `json:"comment"`
	// the event which occured; create, delete, update
	Event string `json:"event"`
}

type ListRequest struct {
}

type ListResponse struct {
	// the comment of comments
	Comments []Comment `json:"comments"`
}

type ReadRequest struct {
	// the comment id
	Id string `json:"id"`
}

type ReadResponse struct {
	// The comment
	Comment *Comment `json:"comment"`
}

type UpdateRequest struct {
	Comment *Comment `json:"comment"`
}

type UpdateResponse struct {
	Comment *Comment `json:"comment"`
}
