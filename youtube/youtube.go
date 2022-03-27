package youtube

import (
	"go.m3o.com/client"
)

type Youtube interface {
	Embed(*EmbedRequest) (*EmbedResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewYoutubeService(token string) *YoutubeService {
	return &YoutubeService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type YoutubeService struct {
	client *client.Client
}

// Embed a YouTube video
func (t *YoutubeService) Embed(request *EmbedRequest) (*EmbedResponse, error) {

	rsp := &EmbedResponse{}
	return rsp, t.client.Call("youtube", "Embed", request, rsp)

}

// Search for videos on YouTube
func (t *YoutubeService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("youtube", "Search", request, rsp)

}

type EmbedRequest struct {
	// provide the youtube url
	Url string `json:"url,omitempty"`
}

type EmbedResponse struct {
	// the embeddable link
	EmbedUrl string `json:"embed_url,omitempty"`
	// the script code
	HtmlScript string `json:"html_script,omitempty"`
	// the full url
	LongUrl string `json:"long_url,omitempty"`
	// the short url
	ShortUrl string `json:"short_url,omitempty"`
}

type SearchRequest struct {
	// Query to search for
	Query string `json:"query,omitempty"`
}

type SearchResponse struct {
	// List of results for the query
	Results []SearchResult `json:"results,omitempty"`
}

type SearchResult struct {
	// if live broadcast then indicates activity:
	// none, upcoming, live, completed
	Broadcasting string `json:"broadcasting,omitempty"`
	// the channel id
	ChannelId string `json:"channel_id,omitempty"`
	// the channel title
	ChannelTitle string `json:"channel_title,omitempty"`
	// the result description
	Description string `json:"description,omitempty"`
	// id of the result
	Id string `json:"id,omitempty"`
	// kind of result: "video", "channel", "playlist"
	Kind string `json:"kind,omitempty"`
	// published at time
	PublishedAt string `json:"published_at,omitempty"`
	// title of the result
	Title string `json:"title,omitempty"`
	// the associated url
	Url string `json:"url,omitempty"`
}
