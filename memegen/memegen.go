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

//
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
	// url of the memegen
	Url string `json:"url"`
}

type Template struct {
	// number of boxes used
	BoxCount int32 `json:"box_count"`
	// height in pixels
	Height int32 `json:"height"`
	// id of the memegen
	Id string `json:"id"`
	// name of the memegen
	Name string `json:"name"`
	// url of the memegen
	Url string `json:"url"`
	// width in pixels
	Width int32 `json:"width"`
}

type TemplatesRequest struct {
}

type TemplatesResponse struct {
	Templates []Template `json:"templates"`
}
