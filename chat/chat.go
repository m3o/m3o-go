package chat

import (
	"go.m3o.com/client"
)

type Chat interface {
	Create(*CreateRequest) (*CreateResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	History(*HistoryRequest) (*HistoryResponse, error)
	Invite(*InviteRequest) (*InviteResponse, error)
	Join(*JoinRequest) (*JoinResponseStream, error)
	Kick(*KickRequest) (*KickResponse, error)
	Leave(*LeaveRequest) (*LeaveResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Send(*SendRequest) (*SendResponse, error)
}

func NewChatService(token string) *ChatService {
	return &ChatService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ChatService struct {
	client *client.Client
}

// Create a new chat room
func (t *ChatService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("chat", "Create", request, rsp)

}

// Delete a chat room
func (t *ChatService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("chat", "Delete", request, rsp)

}

// List the messages in a chat
func (t *ChatService) History(request *HistoryRequest) (*HistoryResponse, error) {

	rsp := &HistoryResponse{}
	return rsp, t.client.Call("chat", "History", request, rsp)

}

// Invite a user to a chat room
func (t *ChatService) Invite(request *InviteRequest) (*InviteResponse, error) {

	rsp := &InviteResponse{}
	return rsp, t.client.Call("chat", "Invite", request, rsp)

}

// Join a chat room
func (t *ChatService) Join(request *JoinRequest) (*JoinResponseStream, error) {
	stream, err := t.client.Stream("chat", "Join", request)
	if err != nil {
		return nil, err
	}
	return &JoinResponseStream{
		stream: stream,
	}, nil

}

type JoinResponseStream struct {
	stream *client.Stream
}

func (t *JoinResponseStream) Recv() (*JoinResponse, error) {
	var rsp JoinResponse
	if err := t.stream.Recv(&rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// Kick a user from a chat room
func (t *ChatService) Kick(request *KickRequest) (*KickResponse, error) {

	rsp := &KickResponse{}
	return rsp, t.client.Call("chat", "Kick", request, rsp)

}

// Leave a chat room
func (t *ChatService) Leave(request *LeaveRequest) (*LeaveResponse, error) {

	rsp := &LeaveResponse{}
	return rsp, t.client.Call("chat", "Leave", request, rsp)

}

// List available chats
func (t *ChatService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("chat", "List", request, rsp)

}

// Connect to a chat to receive a stream of messages
// Send a message to a chat
func (t *ChatService) Send(request *SendRequest) (*SendResponse, error) {

	rsp := &SendResponse{}
	return rsp, t.client.Call("chat", "Send", request, rsp)

}

type CreateRequest struct {
	// chat description
	Description string `json:"description,omitempty"`
	// name of the room
	Name string `json:"name,omitempty"`
	// whether its a private room
	Private bool `json:"private,omitempty"`
	// optional list of user ids
	UserIds []string `json:"user_ids,omitempty"`
}

type CreateResponse struct {
	// the unique chat room
	Room *Room `json:"room,omitempty"`
}

type DeleteRequest struct {
	// the chat room id to delete
	RoomId string `json:"room_id,omitempty"`
}

type DeleteResponse struct {
	Room *Room `json:"room,omitempty"`
}

type HistoryRequest struct {
	// the chat room id to get
	RoomId string `json:"room_id,omitempty"`
}

type HistoryResponse struct {
	// messages in the chat room
	Messages []Message `json:"messages,omitempty"`
}

type InviteRequest struct {
	// the room id
	RoomId string `json:"room_id,omitempty"`
	// the user id
	UserId string `json:"user_id,omitempty"`
}

type InviteResponse struct {
	Room *Room `json:"room,omitempty"`
}

type JoinRequest struct {
	// chat room to join
	RoomId string `json:"room_id,omitempty"`
	// user id joining
	UserId string `json:"user_id,omitempty"`
}

type JoinResponse struct {
	Message *Message `json:"message,omitempty"`
}

type KickRequest struct {
	// the chat room id
	RoomId string `json:"room_id,omitempty"`
	// the user id
	UserId string `json:"user_id,omitempty"`
}

type KickResponse struct {
	Room *Room `json:"room,omitempty"`
}

type LeaveRequest struct {
	// the chat room id
	RoomId string `json:"room_id,omitempty"`
	// the user id
	UserId string `json:"user_id,omitempty"`
}

type LeaveResponse struct {
	Room *Room `json:"room,omitempty"`
}

type ListRequest struct {
	// optional user id to filter by
	UserId string `json:"user_id,omitempty"`
}

type ListResponse struct {
	Rooms []Room `json:"rooms,omitempty"`
}

type Message struct {
	// a client side id, should be validated by the server to make the request retry safe
	Client string `json:"client,omitempty"`
	// id of the message, allocated by the server
	Id string `json:"id,omitempty"`
	// id of the chat the message is being sent to / from
	RoomId string `json:"room_id,omitempty"`
	// time the message was sent in RFC3339 format
	SentAt string `json:"sent_at,omitempty"`
	// subject of the message
	Subject string `json:"subject,omitempty"`
	// text of the message
	Text string `json:"text,omitempty"`
	// id of the user who sent the message
	UserId string `json:"user_id,omitempty"`
}

type Room struct {
	// time of creation
	CreatedAt string `json:"created_at,omitempty"`
	// description of the that
	Description string `json:"description,omitempty"`
	// unique room id
	Id string `json:"id,omitempty"`
	// name of the chat
	Name string `json:"name,omitempty"`
	// whether its a private room
	Private bool `json:"private,omitempty"`
	// list of users
	UserIds []string `json:"user_ids,omitempty"`
}

type SendRequest struct {
	// a client side id, should be validated by the server to make the request retry safe
	Client string `json:"client,omitempty"`
	// id of the chat room the message is being sent to / from
	RoomId string `json:"room_id,omitempty"`
	// subject of the message
	Subject string `json:"subject,omitempty"`
	// text of the message
	Text string `json:"text,omitempty"`
	// id of the user who sent the message
	UserId string `json:"user_id,omitempty"`
}

type SendResponse struct {
	// the message which was created
	Message *Message `json:"message,omitempty"`
}
