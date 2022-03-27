package memegen

import (
	"go.m3o.com/client"
)

type Memegen interface {
	Generate(*GenerateRequest) (*GenerateResponse, error)
	Templates(*TemplatesRequest) (*TemplatesResponse, error)
}

func NewMemegenService(token string) *MemegenService {
	return &MemegenService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type MemegenService struct {
	client *client.Client
}

// Generate a meme using a template
func (t *MemegenService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("memegen", "Generate", request, rsp)

}

// List the available templates
func (t *MemegenService) Templates(request *TemplatesRequest) (*TemplatesResponse, error) {

	rsp := &TemplatesResponse{}
	return rsp, t.client.Call("memegen", "Templates", request, rsp)

}

type Box struct {
	// colour hex code
	Color string `json:"color,omitempty"`
	// height in pixels
	Height int32 `json:"height,omitempty"`
	// outline color hex code
	Outline string `json:"outline,omitempty"`
	// text to display
	Text string `json:"text,omitempty"`
	// width in pixels
	Width int32 `json:"width,omitempty"`
	// x axis position
	X int32 `json:"x,omitempty"`
	// y axis position
	Y int32 `json:"y,omitempty"`
}

type GenerateRequest struct {
	// bottom text
	BottomText string `json:"bottom_text,omitempty"`
	// font: arial or impact
	Font string `json:"font,omitempty"`
	// the template id to use
	Id string `json:"id,omitempty"`
	// font size; defaults to 50px
	MaxFontSize string `json:"max_font_size,omitempty"`
	// top text
	TopText string `json:"top_text,omitempty"`
}

type GenerateResponse struct {
	// url of the memegen
	Url string `json:"url,omitempty"`
}

type Template struct {
	// number of boxes used
	BoxCount int32 `json:"box_count,omitempty"`
	// height in pixels
	Height int32 `json:"height,omitempty"`
	// id of the memegen
	Id string `json:"id,omitempty"`
	// name of the memegen
	Name string `json:"name,omitempty"`
	// url of the memegen
	Url string `json:"url,omitempty"`
	// width in pixels
	Width int32 `json:"width,omitempty"`
}

type TemplatesRequest struct {
}

type TemplatesResponse struct {
	Templates []Template `json:"templates,omitempty"`
}
