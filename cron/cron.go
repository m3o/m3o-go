package cron

import (
	"go.m3o.com/client"
)

type Cron interface {
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Jobs(*JobsRequest) (*JobsResponse, error)
	Schedule(*ScheduleRequest) (*ScheduleResponse, error)
}

func NewCronService(token string) *CronService {
	return &CronService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type CronService struct {
	client *client.Client
}

// Delete a cron job
func (t *CronService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("cron", "Delete", request, rsp)

}

// List all cron jobs
func (t *CronService) Jobs(request *JobsRequest) (*JobsResponse, error) {

	rsp := &JobsResponse{}
	return rsp, t.client.Call("cron", "Jobs", request, rsp)

}

// Schedule a cron job
func (t *CronService) Schedule(request *ScheduleRequest) (*ScheduleResponse, error) {

	rsp := &ScheduleResponse{}
	return rsp, t.client.Call("cron", "Schedule", request, rsp)

}

type DeleteRequest struct {
	// id of the cron job
	Id string `json:"id,omitempty"`
}

type DeleteResponse struct {
}

type Job struct {
	// callback url e.g https://google.com
	Callback string `json:"callback,omitempty"`
	// description
	Description string `json:"description,omitempty"`
	// job id
	Id string `json:"id,omitempty"`
	// scheduled interval
	Interval string `json:"interval,omitempty"`
	// name
	Name string `json:"name,omitempty"`
}

type JobsRequest struct {
}

type JobsResponse struct {
	// the list of scheduled jobs
	Jobs []Job `json:"jobs,omitempty"`
}

type ScheduleRequest struct {
	// callback url e.g https://google.com
	Callback string `json:"callback,omitempty"`
	// description
	Description string `json:"description,omitempty"`
	// unique id of job (optional)
	Id string `json:"id,omitempty"`
	// interval e.g * * * * *
	Interval string `json:"interval,omitempty"`
	// name of cron
	Name string `json:"name,omitempty"`
}

type ScheduleResponse struct {
	// the scheduled job
	Job *Job `json:"job,omitempty"`
}
