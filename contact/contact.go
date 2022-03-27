package contact

import (
	"go.m3o.com/client"
)

type Contact interface {
	Create(*CreateRequest) (*CreateResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Read(*ReadRequest) (*ReadResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewContactService(token string) *ContactService {
	return &ContactService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ContactService struct {
	client *client.Client
}

// Create a contact
func (t *ContactService) Create(request *CreateRequest) (*CreateResponse, error) {

	rsp := &CreateResponse{}
	return rsp, t.client.Call("contact", "Create", request, rsp)

}

// Delete a contact
func (t *ContactService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("contact", "Delete", request, rsp)

}

// List contacts
func (t *ContactService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("contact", "List", request, rsp)

}

// Read contact details
func (t *ContactService) Read(request *ReadRequest) (*ReadResponse, error) {

	rsp := &ReadResponse{}
	return rsp, t.client.Call("contact", "Read", request, rsp)

}

// Update a contact
func (t *ContactService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("contact", "Update", request, rsp)

}

type Address struct {
	// the label of the address
	Label string `json:"label,omitempty"`
	// the address location
	Location string `json:"location,omitempty"`
}

type ContactInfo struct {
	// the address
	Addresses []Address `json:"addresses,omitempty"`
	// the birthday
	Birthday string `json:"birthday,omitempty"`
	// create date string in RFC3339
	CreatedAt string `json:"created_at,omitempty"`
	// the emails
	Emails []Email `json:"emails,omitempty"`
	// contact id
	Id string `json:"id,omitempty"`
	// the contact links
	Links []Link `json:"links,omitempty"`
	// the contact name
	Name string `json:"name,omitempty"`
	// note of the contact
	Note string `json:"note,omitempty"`
	// the phone numbers
	Phones []Phone `json:"phones,omitempty"`
	// the social media username
	SocialMedias *SocialMedia `json:"social_medias,omitempty"`
	// update date string in RFC3339
	UpdatedAt string `json:"updated_at,omitempty"`
}

type CreateRequest struct {
	// optional, location
	Addresses []Address `json:"addresses,omitempty"`
	// optional, birthday
	Birthday string `json:"birthday,omitempty"`
	// optional, emails
	Emails []Email `json:"emails,omitempty"`
	// optional, links
	Links []Link `json:"links,omitempty"`
	// required, the name of the contact
	Name string `json:"name,omitempty"`
	// optional, note of the contact
	Note string `json:"note,omitempty"`
	// optional, phone numbers
	Phones []Phone `json:"phones,omitempty"`
	// optional, social media
	SocialMedias *SocialMedia `json:"social_medias,omitempty"`
}

type CreateResponse struct {
	Contact *ContactInfo `json:"contact,omitempty"`
}

type DeleteRequest struct {
	// the id of the contact
	Id string `json:"id,omitempty"`
}

type DeleteResponse struct {
}

type Email struct {
	// the email address
	Address string `json:"address,omitempty"`
	// the label of the email
	Label string `json:"label,omitempty"`
}

type Link struct {
	// the label of the link
	Label string `json:"label,omitempty"`
	// the url of the contact
	Url string `json:"url,omitempty"`
}

type ListRequest struct {
	// optional, default is 30
	Limit int32 `json:"limit,omitempty"`
	// optional
	Offset int32 `json:"offset,omitempty"`
}

type ListResponse struct {
	Contacts []ContactInfo `json:"contacts,omitempty"`
}

type Phone struct {
	// the label of the phone number
	Label string `json:"label,omitempty"`
	// phone number
	Number string `json:"number,omitempty"`
}

type ReadRequest struct {
	Id string `json:"id,omitempty"`
}

type ReadResponse struct {
	Contact *ContactInfo `json:"contact,omitempty"`
}

type SocialMedia struct {
	// the label of the social
	Label string `json:"label,omitempty"`
	// the username of social media
	Username string `json:"username,omitempty"`
}

type UpdateRequest struct {
	// optional, addresses
	Addresses []Address `json:"addresses,omitempty"`
	// optional, birthday
	Birthday string `json:"birthday,omitempty"`
	// optional, emails
	Emails []Email `json:"emails,omitempty"`
	// required, the contact id
	Id string `json:"id,omitempty"`
	// optional, links
	Links []Link `json:"links,omitempty"`
	// required, the name
	Name string `json:"name,omitempty"`
	// optional, note
	Note string `json:"note,omitempty"`
	// optional, phone number
	Phones []Phone `json:"phones,omitempty"`
	// optional, social media
	SocialMedias *SocialMedia `json:"social_medias,omitempty"`
}

type UpdateResponse struct {
	Contact *ContactInfo `json:"contact,omitempty"`
}
