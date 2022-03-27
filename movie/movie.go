package movie

import (
	"go.m3o.com/client"
)

type Movie interface {
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewMovieService(token string) *MovieService {
	return &MovieService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type MovieService struct {
	client *client.Client
}

// Search for movies by simple text search
func (t *MovieService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("movie", "Search", request, rsp)

}

type MovieInfo struct {
	Adult            bool    `json:"adult,omitempty"`
	BackdropPath     string  `json:"backdrop_path,omitempty"`
	GenreIds         int32   `json:"genre_ids,omitempty"`
	Id               int32   `json:"id,omitempty"`
	OriginalLanguage string  `json:"original_language,omitempty"`
	OriginalTitle    string  `json:"original_title,omitempty"`
	Overview         string  `json:"overview,omitempty"`
	Popularity       float64 `json:"popularity,omitempty"`
	PosterPath       string  `json:"poster_path,omitempty"`
	ReleaseDate      string  `json:"release_date,omitempty"`
	Title            string  `json:"title,omitempty"`
	Video            bool    `json:"video,omitempty"`
	VoteAverage      float64 `json:"vote_average,omitempty"`
	VoteCount        int32   `json:"vote_count,omitempty"`
}

type SearchRequest struct {
	// a ISO 639-1 value to display translated data
	Language string `json:"language,omitempty"`
	// page to query
	Page int32 `json:"page,omitempty"`
	// year of release
	PrimaryReleaseYear int32 `json:"primary_release_year,omitempty"`
	// a text query to search
	Query string `json:"query,omitempty"`
	// a ISO 3166-1 code to filter release dates.
	Region string `json:"region,omitempty"`
	// year of making
	Year int32 `json:"year,omitempty"`
}

type SearchResponse struct {
	Page         int32       `json:"page,omitempty"`
	Results      []MovieInfo `json:"results,omitempty"`
	TotalPages   int32       `json:"total_pages,omitempty"`
	TotalResults int32       `json:"total_results,omitempty"`
}
