package gifs

import (
	"go.m3o.com/client"
)

type Gifs interface {
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewGifsService(token string) *GifsService {
	return &GifsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type GifsService struct {
	client *client.Client
}

// Search for a GIF
func (t *GifsService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("gifs", "Search", request, rsp)

}

type Gif struct {
	// URL used for embedding the GIF
	EmbedUrl string `json:"embed_url,omitempty"`
	// The ID of the GIF
	Id string `json:"id,omitempty"`
	// The different formats available for this GIF
	Images *ImageFormats `json:"images,omitempty"`
	// The content rating for the GIF
	Rating string `json:"rating,omitempty"`
	// A short URL for this GIF
	ShortUrl string `json:"short_url,omitempty"`
	// The slug used in the GIF's URL
	Slug string `json:"slug,omitempty"`
	// The page on which this GIF was found
	Source string `json:"source,omitempty"`
	// The title for this GIF
	Title string `json:"title,omitempty"`
	// The URL for this GIF
	Url string `json:"url,omitempty"`
}

type ImageFormat struct {
	// height
	Height int32 `json:"height,omitempty"`
	// size of the MP4 version
	Mp4Size int32 `json:"mp4_size,omitempty"`
	// URL to an MP4 version of the gif
	Mp4Url string `json:"mp4_url,omitempty"`
	// size in bytes
	Size int32 `json:"size,omitempty"`
	// URL of the gif
	Url string `json:"url,omitempty"`
	// size of the webp version
	WebpSize int32 `json:"webp_size,omitempty"`
	// URL to a webp version of the gif
	WebpUrl string `json:"webp_url,omitempty"`
	// width
	Width int32 `json:"width,omitempty"`
}

type ImageFormats struct {
	// A downsized version of the GIF < 2MB
	Downsized *ImageFormat `json:"downsized,omitempty"`
	// A downsized version of the GIF < 8MB
	DownsizedLarge *ImageFormat `json:"downsized_large,omitempty"`
	// A downsized version of the GIF < 5MB
	DownsizedMedium *ImageFormat `json:"downsized_medium,omitempty"`
	// A downsized version of the GIF < 200kb
	DownsizedSmall *ImageFormat `json:"downsized_small,omitempty"`
	// Static image of the downsized version of the GIF
	DownsizedStill *ImageFormat `json:"downsized_still,omitempty"`
	// Version of the GIF with fixed height of 200 pixels. Good for mobile use
	FixedHeight *ImageFormat `json:"fixed_height,omitempty"`
	// Version of the GIF with fixed height of 200 pixels and number of frames reduced to 6
	FixedHeightDownsampled *ImageFormat `json:"fixed_height_downsampled,omitempty"`
	// Version of the GIF with fixed height of 100 pixels. Good for mobile keyboards
	FixedHeightSmall *ImageFormat `json:"fixed_height_small,omitempty"`
	// Static image of the GIF with fixed height of 100 pixels
	FixedHeightSmallStill *ImageFormat `json:"fixed_height_small_still,omitempty"`
	// Static image of the GIF with fixed height of 200 pixels
	FixedHeightStill *ImageFormat `json:"fixed_height_still,omitempty"`
	// Version of the GIF with fixed width of 200 pixels. Good for mobile use
	FixedWidth *ImageFormat `json:"fixed_width,omitempty"`
	// Version of the GIF with fixed width of 200 pixels and number of frames reduced to 6
	FixedWidthDownsampled *ImageFormat `json:"fixed_width_downsampled,omitempty"`
	// Version of the GIF with fixed width of 100 pixels. Good for mobile keyboards
	FixedWidthSmall *ImageFormat `json:"fixed_width_small,omitempty"`
	// Static image of the GIF with fixed width of 100 pixels
	FixedWidthSmallStill *ImageFormat `json:"fixed_width_small_still,omitempty"`
	// Static image of the GIF with fixed width of 200 pixels
	FixedWidthStill *ImageFormat `json:"fixed_width_still,omitempty"`
	// 15 second version of the GIF looping
	Looping *ImageFormat `json:"looping,omitempty"`
	// The original GIF. Good for desktop use
	Original *ImageFormat `json:"original,omitempty"`
	// Static image of the original version of the GIF
	OriginalStill *ImageFormat `json:"original_still,omitempty"`
	// mp4 version of the GIF <50kb displaying first 1-2 secs
	Preview *ImageFormat `json:"preview,omitempty"`
	// Version of the GIF <50kb displaying first 1-2 secs
	PreviewGif *ImageFormat `json:"preview_gif,omitempty"`
}

type Pagination struct {
	// total number returned in this response
	Count int32 `json:"count,omitempty"`
	// position in pagination
	Offset int32 `json:"offset,omitempty"`
	// total number of results available
	TotalCount int32 `json:"total_count,omitempty"`
}

type SearchRequest struct {
	// ISO 2 letter language code for regional content
	Lang string `json:"lang,omitempty"`
	// Max number of gifs to return. Defaults to 25
	Limit int32 `json:"limit,omitempty"`
	// The start position of results (used with pagination)
	Offset int32 `json:"offset,omitempty"`
	// The search term
	Query string `json:"query,omitempty"`
	// Apply age related content filter. "g", "pg", "pg-13", or "r". Defaults to "g"
	Rating string `json:"rating,omitempty"`
}

type SearchResponse struct {
	// list of results
	Data []Gif `json:"data,omitempty"`
	// information on pagination
	Pagination *Pagination `json:"pagination,omitempty"`
}
