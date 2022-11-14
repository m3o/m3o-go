package dns

import (
	"go.m3o.com/client"
)

type Dns interface {
	Query(*QueryRequest) (*QueryResponse, error)
	Whois(*WhoisRequest) (*WhoisResponse, error)
}

func NewDnsService(token string) *DnsService {
	return &DnsService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type DnsService struct {
	client *client.Client
}

// Query an addresss
func (t *DnsService) Query(request *QueryRequest) (*QueryResponse, error) {

	rsp := &QueryResponse{}
	return rsp, t.client.Call("dns", "Query", request, rsp)

}

// Check who owns a domain
func (t *DnsService) Whois(request *WhoisRequest) (*WhoisResponse, error) {

	rsp := &WhoisResponse{}
	return rsp, t.client.Call("dns", "Whois", request, rsp)

}

type Answer struct {
	// time to live
	Ttl int32 `json:"TTL,omitempty"`
	// the answer
	Data string `json:"data,omitempty"`
	// name resolved
	Name string `json:"name,omitempty"`
	// type of record
	Type int32 `json:"type,omitempty"`
}

type Domain struct {
	// domain id
	Id string `json:"id,omitempty"`
}

type QueryRequest struct {
	// name to resolve
	Name string `json:"name,omitempty"`
	// type of query e.g A, AAAA, MX, SRV
	Type string `json:"type,omitempty"`
}

type QueryResponse struct {
	Ad       bool       `json:"AD,omitempty"`
	Cd       bool       `json:"CD,omitempty"`
	Ra       bool       `json:"RA,omitempty"`
	Rd       bool       `json:"RD,omitempty"`
	Tc       bool       `json:"TC,omitempty"`
	Answer   []Answer   `json:"answer,omitempty"`
	Provider string     `json:"provider,omitempty"`
	Question []Question `json:"question,omitempty"`
	Status   int32      `json:"status,omitempty"`
}

type Question struct {
	// name to query
	Name string `json:"name,omitempty"`
	// type of record
	Type int32 `json:"type,omitempty"`
}

type WhoisRequest struct {
	Domain string `json:"domain,omitempty"`
}

type WhoisResponse struct {
	// abuse email
	AbuseEmail string `json:"abuse_email,omitempty"`
	// abuse phone
	AbusePhone string `json:"abuse_phone,omitempty"`
	// time of creation
	Created string `json:"created,omitempty"`
	// domain name
	Domain string `json:"domain,omitempty"`
	// time of expiry
	Expiry string `json:"expiry,omitempty"`
	// domain id
	Id string `json:"id,omitempty"`
	// nameservers
	Nameservers []string `json:"nameservers,omitempty"`
	// the registrar
	Registrar string `json:"registrar,omitempty"`
	// the registrar iana id
	RegistrarId string `json:"registrar_id,omitempty"`
	// registrar
	RegistrarUrl string `json:"registrar_url,omitempty"`
	// status of domain
	Status []string `json:"status,omitempty"`
	// time of update
	Updated string `json:"updated,omitempty"`
	// whois server
	WhoisServer string `json:"whois_server,omitempty"`
}
