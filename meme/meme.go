package meme

import (
	"go.m3o.com/client"
)

type Meme interface {
	Generate(*GenerateRequest) (*GenerateResponse, error)
	Templates(*TemplatesRequest) (*TemplatesResponse, error)
}

func NewMemeService(token string) *MemeService {
	return &MemeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type MemeService struct {
	client *client.Client
}

//
func (t *MemeService) Generate(request *GenerateRequest) (*GenerateResponse, error) {

	rsp := &GenerateResponse{}
	return rsp, t.client.Call("meme", "Generate", request, rsp)

}

// List the available templates
func (t *MemeService) Templates(request *TemplatesRequest) (*TemplatesResponse, error) {

	rsp := &TemplatesResponse{}
	return rsp, t.client.Call("meme", "Templates", request, rsp)

}

type Box struct {
	// colour hex code
	Color string `json:"color"`
	// height in pixels
	Height int32 `json:"height"`
	// outline color hex code
	Outline string `json:"outline"`
	// text to display
	Text string `json:"text"`
	// width in pixels
	Width int32 `json:"width"`
	// x axis position
	X int32 `json:"x"`
	// y axis position
	Y int32 `json:"y"`
}

type GenerateRequest struct {
	// bottom text
	BottomText string `json:"bottom_text"`
	// font: arial or impact
	Font string `json:"font"`
	// the template id to use
	Id string `json:"id"`
	// font size; defaults to 50px
	MaxFontSize string `json:"max_font_size"`
	// top text
	TopText string `json:"top_text"`
}

type GenerateResponse struct {
	// url of the meme
	Url string `json:"url"`
}

type Template struct {
	// number of boxes used
	BoxCount int32 `json:"box_count"`
	// height in pixels
	Height int32 `json:"height"`
	// id of the meme
	Id string `json:"id"`
	// name of the meme
	Name string `json:"name"`
	// url of the meme
	Url string `json:"url"`
	// width in pixels
	Width int32 `json:"width"`
}

type TemplatesRequest struct {
}

type TemplatesResponse struct {
	Templates []Template `json:"templates"`
}
