package app

import (
	"go.m3o.com/client"
)

type App interface {
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Get(*GetRequest) (*GetResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Logs(*LogsRequest) (*LogsResponse, error)
	Regions(*RegionsRequest) (*RegionsResponse, error)
	Reserve(*ReserveRequest) (*ReserveResponse, error)
	Resolve(*ResolveRequest) (*ResolveResponse, error)
	Run(*RunRequest) (*RunResponse, error)
	Status(*StatusRequest) (*StatusResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewAppService(token string) *AppService {
	return &AppService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AppService struct {
	client *client.Client
}

// Delete an app
func (t *AppService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("app", "Delete", request, rsp)

}

// Get an app by name
func (t *AppService) Get(request *GetRequest) (*GetResponse, error) {

	rsp := &GetResponse{}
	return rsp, t.client.Call("app", "Get", request, rsp)

}

// List all the apps
func (t *AppService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("app", "List", request, rsp)

}

// Get the logs for an app
func (t *AppService) Logs(request *LogsRequest) (*LogsResponse, error) {

	rsp := &LogsResponse{}
	return rsp, t.client.Call("app", "Logs", request, rsp)

}

// Return the support regions
func (t *AppService) Regions(request *RegionsRequest) (*RegionsResponse, error) {

	rsp := &RegionsResponse{}
	return rsp, t.client.Call("app", "Regions", request, rsp)

}

// Reserve app names
func (t *AppService) Reserve(request *ReserveRequest) (*ReserveResponse, error) {

	rsp := &ReserveResponse{}
	return rsp, t.client.Call("app", "Reserve", request, rsp)

}

// Resolve an app by id to its raw backend endpoint
func (t *AppService) Resolve(request *ResolveRequest) (*ResolveResponse, error) {

	rsp := &ResolveResponse{}
	return rsp, t.client.Call("app", "Resolve", request, rsp)

}

// Run an app from source
func (t *AppService) Run(request *RunRequest) (*RunResponse, error) {

	rsp := &RunResponse{}
	return rsp, t.client.Call("app", "Run", request, rsp)

}

// Get the status of an app
func (t *AppService) Status(request *StatusRequest) (*StatusResponse, error) {

	rsp := &StatusResponse{}
	return rsp, t.client.Call("app", "Status", request, rsp)

}

// Update the app. The latest source code will be downloaded, built and deployed.
func (t *AppService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("app", "Update", request, rsp)

}

type DeleteRequest struct {
	// name of the app
	Name string `json:"name,omitempty"`
}

type DeleteResponse struct {
}

type GetRequest struct {
	Name string `json:"name,omitempty"`
}

type GetResponse struct {
	Service *Service `json:"service,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	// all the apps
	Services []Service `json:"services,omitempty"`
}

type LogsRequest struct {
	// type of logs to retrieve, currently supported options - "build"
	LogsType string `json:"logs_type,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
}

type LogsResponse struct {
	Logs string `json:"logs,omitempty"`
}

type RegionsRequest struct {
}

type RegionsResponse struct {
	Regions []string `json:"regions,omitempty"`
}

type Reservation struct {
	// time of reservation
	Created string `json:"created,omitempty"`
	// time reservation expires
	Expires string `json:"expires,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
	// owner id
	Owner string `json:"owner,omitempty"`
	// associated token
	Token string `json:"token,omitempty"`
}

type ReserveRequest struct {
	// name of your app e.g helloworld
	Name string `json:"name,omitempty"`
}

type ReserveResponse struct {
	// The app reservation
	Reservation *Reservation `json:"reservation,omitempty"`
}

type ResolveRequest struct {
	// the service id
	Id string `json:"id,omitempty"`
}

type ResolveResponse struct {
	// the end provider url
	Url string `json:"url,omitempty"`
}

type RunRequest struct {
	// branch. defaults to master
	Branch string `json:"branch,omitempty"`
	// associated env vars to pass in
	EnvVars map[string]string `json:"env_vars,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
	// port to run on
	Port int32 `json:"port,omitempty"`
	// region to run in
	Region string `json:"region,omitempty"`
	// source repository
	Repo string `json:"repo,omitempty"`
}

type RunResponse struct {
	// The running service
	Service *Service `json:"service,omitempty"`
}

type Service struct {
	// raw backend endpoint
	Backend string `json:"backend,omitempty"`
	// branch of code
	Branch string `json:"branch,omitempty"`
	// time of creation
	Created string `json:"created,omitempty"`
	// custom domains
	CustomDomains []string `json:"custom_domains,omitempty"`
	// associated env vars
	EnvVars map[string]string `json:"env_vars,omitempty"`
	// unique id
	Id string `json:"id,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
	// port running on
	Port int32 `json:"port,omitempty"`
	// region running in
	Region string `json:"region,omitempty"`
	// source repository
	Repo string `json:"repo,omitempty"`
	// status of the app
	Status string `json:"status,omitempty"`
	// last updated
	Updated string `json:"updated,omitempty"`
	// app url
	Url string `json:"url,omitempty"`
}

type StatusRequest struct {
	// name of the app
	Name string `json:"name,omitempty"`
}

type StatusResponse struct {
	// running service info
	Service *Service `json:"service,omitempty"`
}

type UpdateRequest struct {
	// Additional env vars to update
	EnvVars map[string]string `json:"env_vars,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
}

type UpdateResponse struct {
}
