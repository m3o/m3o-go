package twitter

import (
	"go.m3o.com/client"
)

type Twitter interface {
	Search(*SearchRequest) (*SearchResponse, error)
	Timeline(*TimelineRequest) (*TimelineResponse, error)
	Trends(*TrendsRequest) (*TrendsResponse, error)
	User(*UserRequest) (*UserResponse, error)
}

func NewTwitterService(token string) *TwitterService {
	return &TwitterService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type TwitterService struct {
	client *client.Client
}

// Search for tweets with a simple query
func (t *TwitterService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("twitter", "Search", request, rsp)

}

// Get the timeline for a given user
func (t *TwitterService) Timeline(request *TimelineRequest) (*TimelineResponse, error) {

	rsp := &TimelineResponse{}
	return rsp, t.client.Call("twitter", "Timeline", request, rsp)

}

// Get the current global trending topics
func (t *TwitterService) Trends(request *TrendsRequest) (*TrendsResponse, error) {

	rsp := &TrendsResponse{}
	return rsp, t.client.Call("twitter", "Trends", request, rsp)

}

// Get a user's twitter profile
func (t *TwitterService) User(request *UserRequest) (*UserResponse, error) {

	rsp := &UserResponse{}
	return rsp, t.client.Call("twitter", "User", request, rsp)

}

type Profile struct {
	// the account creation date
	CreatedAt string `json:"created_at,omitempty"`
	// the user description
	Description string `json:"description,omitempty"`
	// the follower count
	Followers int64 `json:"followers,string,omitempty"`
	// the user id
	Id int64 `json:"id,string,omitempty"`
	// The user's profile picture
	ImageUrl string `json:"image_url,omitempty"`
	// the user's location
	Location string `json:"location,omitempty"`
	// display name of the user
	Name string `json:"name,omitempty"`
	// if the account is private
	Private bool `json:"private,omitempty"`
	// the username
	Username string `json:"username,omitempty"`
	// if the account is verified
	Verified bool `json:"verified,omitempty"`
}

type SearchRequest struct {
	// number of tweets to return. default: 20
	Limit int32 `json:"limit,omitempty"`
	// the query to search for
	Query string `json:"query,omitempty"`
}

type SearchResponse struct {
	// the related tweets for the search
	Tweets []Tweet `json:"tweets,omitempty"`
}

type TimelineRequest struct {
	// number of tweets to return. default: 20
	Limit int32 `json:"limit,omitempty"`
	// the username to request the timeline for
	Username string `json:"username,omitempty"`
}

type TimelineResponse struct {
	// The recent tweets for the user
	Tweets []Tweet `json:"tweets,omitempty"`
}

type Trend struct {
	// name of the trend
	Name string `json:"name,omitempty"`
	// the volume of tweets in last 24 hours
	TweetVolume int64 `json:"tweet_volume,string,omitempty"`
	// the twitter url
	Url string `json:"url,omitempty"`
}

type TrendsRequest struct {
}

type TrendsResponse struct {
	// a list of trending topics
	Trends []Trend `json:"trends,omitempty"`
}

type Tweet struct {
	// time of tweet
	CreatedAt string `json:"created_at,omitempty"`
	// number of times favourited
	FavouritedCount int64 `json:"favourited_count,string,omitempty"`
	// id of the tweet
	Id int64 `json:"id,string,omitempty"`
	// number of times retweeted
	RetweetedCount int64 `json:"retweeted_count,string,omitempty"`
	// text of the tweet
	Text string `json:"text,omitempty"`
	// username of the person who tweeted
	Username string `json:"username,omitempty"`
}

type UserRequest struct {
	// the username to lookup
	Username string `json:"username,omitempty"`
}

type UserResponse struct {
	// The requested user profile
	Profile *Profile `json:"profile,omitempty"`
	// the current user status
	Status *Tweet `json:"status,omitempty"`
}
