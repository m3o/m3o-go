package search

import (
	"go.m3o.com/client"
)

type Search interface {
	CreateIndex(*CreateIndexRequest) (*CreateIndexResponse, error)
	DeleteIndex(*DeleteIndexRequest) (*DeleteIndexResponse, error)
	Delete(*DeleteRequest) (*DeleteResponse, error)
	Index(*IndexRequest) (*IndexResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewSearchService(token string) *SearchService {
	return &SearchService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type SearchService struct {
	client *client.Client
}

// Create an index by name
func (t *SearchService) CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error) {

	rsp := &CreateIndexResponse{}
	return rsp, t.client.Call("search", "CreateIndex", request, rsp)

}

// Delete an index by name
func (t *SearchService) DeleteIndex(request *DeleteIndexRequest) (*DeleteIndexResponse, error) {

	rsp := &DeleteIndexResponse{}
	return rsp, t.client.Call("search", "DeleteIndex", request, rsp)

}

// Delete a record given its ID
func (t *SearchService) Delete(request *DeleteRequest) (*DeleteResponse, error) {

	rsp := &DeleteResponse{}
	return rsp, t.client.Call("search", "Delete", request, rsp)

}

// Index a record i.e. insert a document to search for.
func (t *SearchService) Index(request *IndexRequest) (*IndexResponse, error) {

	rsp := &IndexResponse{}
	return rsp, t.client.Call("search", "Index", request, rsp)

}

// Search for records in a given in index
func (t *SearchService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("search", "Search", request, rsp)

}

type CreateIndexRequest struct {
	// The name of the index
	Index string `json:"index,omitempty"`
}

type CreateIndexResponse struct {
}

type DeleteIndexRequest struct {
	// The name of the index to delete
	Index string `json:"index,omitempty"`
}

type DeleteIndexResponse struct {
}

type DeleteRequest struct {
	// The ID of the record to delete
	Id string `json:"id,omitempty"`
	// The index the record belongs to
	Index string `json:"index,omitempty"`
}

type DeleteResponse struct {
}

type Field struct {
	// The name of the field. Use a `.` separator to define nested fields e.g. foo.bar
	Name string `json:"name,omitempty"`
	// The type of the field - string, number
	Type string `json:"type,omitempty"`
}

type IndexRequest struct {
	// The data to index
	Data map[string]interface{} `json:"data,omitempty"`
	// Optional ID for the record
	Id string `json:"id,omitempty"`
	// The index this record belongs to
	Index string `json:"index,omitempty"`
}

type IndexResponse struct {
	// the indexed record
	Record *Record `json:"record,omitempty"`
}

type Record struct {
	// The JSON contents of the record
	Data map[string]interface{} `json:"data,omitempty"`
	// The ID for this record. If blank, one will be generated
	Id string `json:"id,omitempty"`
}

type SearchRequest struct {
	// The index the record belongs to
	Index string `json:"index,omitempty"`
	// The query. See docs for query language examples
	Query string `json:"query,omitempty"`
}

type SearchResponse struct {
	// The matching records
	Records []Record `json:"records,omitempty"`
}
