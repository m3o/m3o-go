package function

import (
	"go.m3o.com/client"
)

type Function interface {
	Call(*CallRequest) (*CallResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Deploy(*DeployRequest) (*DeployResponse, error)
	Describe(*DescribeRequest) (*DescribeResponse, error)
	List(*ListRequest) (*ListResponse, error)
	Logs(*LogsRequest) (*LogsResponse, error)
	Proxy(*ProxyRequest) (*ProxyResponse, error)
	Regions(*RegionsRequest) (*RegionsResponse, error)
	Reserve(*ReserveRequest) (*ReserveResponse, error)
	Runtimes(*RuntimesRequest) (*RuntimesResponse, error)
	Update(*UpdateRequest) (*UpdateResponse, error)
}

func NewFunctionService(token string) *FunctionService {
	return &FunctionService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type FunctionService struct {
	client *client.Client
}

// Call a function by name
func (t *FunctionService) Call(request *CallRequest) (*CallResponse, error) {

	rsp := &CallResponse{}
	return rsp, t.client.Call("function", "Call", request, rsp)

}

// Delete a function by name
func (t *FunctionService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("function", "Delete", request, rsp)

}

// Deploy a group of functions
func (t *FunctionService) Deploy(request *DeployRequest) (*DeployResponse, error) {

	rsp := &DeployResponse{}
	return rsp, t.client.Call("function", "Deploy", request, rsp)

}

// Get the info for a deployed function
func (t *FunctionService) Describe(request *DescribeRequest) (*DescribeResponse, error) {

	rsp := &DescribeResponse{}
	return rsp, t.client.Call("function", "Describe", request, rsp)

}

// List all the deployed functions
func (t *FunctionService) List(request *ListRequest) (*ListResponse, error) {

	rsp := &ListResponse{}
	return rsp, t.client.Call("function", "List", request, rsp)

}

// Get the logs for a function
func (t *FunctionService) Logs(request *LogsRequest) (*LogsResponse, error) {

	rsp := &LogsResponse{}
	return rsp, t.client.Call("function", "Logs", request, rsp)

}

// Return the backend url for proxying
func (t *FunctionService) Proxy(request *ProxyRequest) (*ProxyResponse, error) {

	rsp := &ProxyResponse{}
	return rsp, t.client.Call("function", "Proxy", request, rsp)

}

// Return a list of supported regions
func (t *FunctionService) Regions(request *RegionsRequest) (*RegionsResponse, error) {

	rsp := &RegionsResponse{}
	return rsp, t.client.Call("function", "Regions", request, rsp)

}

// Reserve function names and resources beyond free quota
func (t *FunctionService) Reserve(request *ReserveRequest) (*ReserveResponse, error) {

	rsp := &ReserveResponse{}
	return rsp, t.client.Call("function", "Reserve", request, rsp)

}

// Return a list of supported runtimes
func (t *FunctionService) Runtimes(request *RuntimesRequest) (*RuntimesResponse, error) {

	rsp := &RuntimesResponse{}
	return rsp, t.client.Call("function", "Runtimes", request, rsp)

}

// Update a function. Downloads the source, builds and redeploys
func (t *FunctionService) Update(request *UpdateRequest) (*UpdateResponse, error) {

	rsp := &UpdateResponse{}
	return rsp, t.client.Call("function", "Update", request, rsp)

}

type CallRequest struct {
	// Name of the function
	Name string `json:"name,omitempty"`
	// Request body that will be passed to the function
	Request map[string]interface{} `json:"request,omitempty"`
}

type CallResponse struct {
	// Response body that the function returned
	Response map[string]interface{} `json:"response,omitempty"`
}

type DeleteRequest struct {
	// The name of the function
	Name string `json:"name,omitempty"`
}

type DeleteResponse struct {
}

type DeployRequest struct {
	// branch to deploy. defaults to master
	Branch string `json:"branch,omitempty"`
	// entry point, ie. handler name in the source code
	// if not provided, defaults to the name parameter
	Entrypoint string `json:"entrypoint,omitempty"`
	// environment variables to pass in at runtime
	EnvVars map[string]string `json:"env_vars,omitempty"`
	// function name
	Name string `json:"name,omitempty"`
	// region to deploy in. defaults to europe-west1
	Region string `json:"region,omitempty"`
	// github url for a repo
	Repo string `json:"repo,omitempty"`
	// runtime/lanaguage of the function e.g php74,
	// nodejs6, nodejs8, nodejs10, nodejs12, nodejs14, nodejs16,
	// dotnet3, java11, ruby26, ruby27, go111, go113, go116,
	// python37, python38, python39
	Runtime string `json:"runtime,omitempty"`
	// inline source code
	Source string `json:"source,omitempty"`
	// optional subfolder path
	Subfolder string `json:"subfolder,omitempty"`
}

type DeployResponse struct {
	Function *Func `json:"function,omitempty"`
}

type DescribeRequest struct {
	// The name of the function
	Name string `json:"name,omitempty"`
}

type DescribeResponse struct {
	// The function requested
	Function *Func `json:"function,omitempty"`
}

type Func struct {
	// branch to deploy. defaults to master
	Branch string `json:"branch,omitempty"`
	// time of creation
	Created string `json:"created,omitempty"`
	// name of handler in source code
	Entrypoint string `json:"entrypoint,omitempty"`
	// associated env vars
	EnvVars map[string]string `json:"env_vars,omitempty"`
	// id of the function
	Id string `json:"id,omitempty"`
	// function name
	// limitation: must be unique across projects
	Name string `json:"name,omitempty"`
	// region to deploy in. defaults to europe-west1
	Region string `json:"region,omitempty"`
	// git repo address
	Repo string `json:"repo,omitempty"`
	// runtime/language of the function e.g php74,
	// nodejs6, nodejs8, nodejs10, nodejs12, nodejs14, nodejs16,
	// dotnet3, java11, ruby26, ruby27, go111, go113, go116,
	// python37, python38, python39
	Runtime string `json:"runtime,omitempty"`
	// the source code
	Source string `json:"source,omitempty"`
	// eg. ACTIVE, DEPLOY_IN_PROGRESS, OFFLINE etc
	Status string `json:"status,omitempty"`
	// subfolder path to entrypoint
	Subfolder string `json:"subfolder,omitempty"`
	// time it was updated
	Updated string `json:"updated,omitempty"`
	// unique url of the function
	Url string `json:"url,omitempty"`
}

type ListRequest struct {
}

type ListResponse struct {
	// List of functions deployed
	Functions []Func `json:"functions,omitempty"`
}

type LogsRequest struct {
	// type of logs to retrieve, currently supported options - "build"
	LogsType string `json:"logs_type,omitempty"`
	// the name of the function
	Name string `json:"name,omitempty"`
}

type LogsResponse struct {
	Logs string `json:"logs,omitempty"`
}

type ProxyRequest struct {
	// id of the function
	Id string `json:"id,omitempty"`
}

type ProxyResponse struct {
	// backend url
	Url string `json:"url,omitempty"`
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

type RuntimesRequest struct {
}

type RuntimesResponse struct {
	Runtimes []string `json:"runtimes,omitempty"`
}

type UpdateRequest struct {
	// function name
	Name string `json:"name,omitempty"`
	// inline source code
	Source string `json:"source,omitempty"`
}

type UpdateResponse struct {
}
