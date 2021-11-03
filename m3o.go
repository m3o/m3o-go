package m3o

import (
	"go.m3o.com/notes"
)

func NewClient(token string) *Client {
	return &Client{
		token: token,

		NotesService: notes.NewNotesService(token),
	}
}

type Client struct {
	token string

	NotesService *notes.NotesService
}
