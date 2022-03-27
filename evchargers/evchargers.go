package evchargers

import (
	"go.m3o.com/client"
)

type Evchargers interface {
	ReferenceData(*ReferenceDataRequest) (*ReferenceDataResponse, error)
	Search(*SearchRequest) (*SearchResponse, error)
}

func NewEvchargersService(token string) *EvchargersService {
	return &EvchargersService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type EvchargersService struct {
	client *client.Client
}

// Retrieve reference data as used by this API and in conjunction with the Search endpoint
func (t *EvchargersService) ReferenceData(request *ReferenceDataRequest) (*ReferenceDataResponse, error) {

	rsp := &ReferenceDataResponse{}
	return rsp, t.client.Call("evchargers", "ReferenceData", request, rsp)

}

// Search by giving a coordinate and a max distance, or bounding box and optional filters
func (t *EvchargersService) Search(request *SearchRequest) (*SearchResponse, error) {

	rsp := &SearchResponse{}
	return rsp, t.client.Call("evchargers", "Search", request, rsp)

}

type Address struct {
	// Any comments about how to access the charger
	AccessComments  string       `json:"access_comments,omitempty"`
	AddressLine1    string       `json:"address_line_1,omitempty"`
	AddressLine2    string       `json:"address_line_2,omitempty"`
	Country         *Country     `json:"country,omitempty"`
	CountryId       string       `json:"country_id,omitempty"`
	LatLng          string       `json:"lat_lng,omitempty"`
	Location        *Coordinates `json:"location,omitempty"`
	Postcode        string       `json:"postcode,omitempty"`
	StateOrProvince string       `json:"state_or_province,omitempty"`
	Title           string       `json:"title,omitempty"`
	Town            string       `json:"town,omitempty"`
}

type BoundingBox struct {
	BottomLeft *Coordinates `json:"bottom_left,omitempty"`
	TopRight   *Coordinates `json:"top_right,omitempty"`
}

type ChargerType struct {
	Comments string `json:"comments,omitempty"`
	Id       string `json:"id,omitempty"`
	// Is this 40KW+
	IsFastChargeCapable bool   `json:"is_fast_charge_capable,omitempty"`
	Title               string `json:"title,omitempty"`
}

type CheckinStatusType struct {
	Id          string `json:"id,omitempty"`
	IsAutomated bool   `json:"is_automated,omitempty"`
	IsPositive  bool   `json:"is_positive,omitempty"`
	Title       string `json:"title,omitempty"`
}

type Connection struct {
	// The amps offered
	Amps           float64         `json:"amps,omitempty"`
	ConnectionType *ConnectionType `json:"connection_type,omitempty"`
	// The ID of the connection type
	ConnectionTypeId string `json:"connection_type_id,omitempty"`
	// The current
	Current string       `json:"current,omitempty"`
	Level   *ChargerType `json:"level,omitempty"`
	// The level of charging power available
	LevelId string `json:"level_id,omitempty"`
	// The power in KW
	Power     float64 `json:"power,omitempty"`
	Reference string  `json:"reference,omitempty"`
	// The voltage offered
	Voltage float64 `json:"voltage,omitempty"`
}

type ConnectionType struct {
	FormalName     string `json:"formal_name,omitempty"`
	Id             string `json:"id,omitempty"`
	IsDiscontinued bool   `json:"is_discontinued,omitempty"`
	IsObsolete     bool   `json:"is_obsolete,omitempty"`
	Title          string `json:"title,omitempty"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Country struct {
	ContinentCode string `json:"continent_code,omitempty"`
	Id            string `json:"id,omitempty"`
	IsoCode       string `json:"iso_code,omitempty"`
	Title         string `json:"title,omitempty"`
}

type CurrentType struct {
	Description string `json:"description,omitempty"`
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
}

type DataProvider struct {
	Comments               string                  `json:"comments,omitempty"`
	DataProviderStatusType *DataProviderStatusType `json:"data_provider_status_type,omitempty"`
	Id                     string                  `json:"id,omitempty"`
	// How is this data licensed
	License string `json:"license,omitempty"`
	Title   string `json:"title,omitempty"`
	Website string `json:"website,omitempty"`
}

type DataProviderStatusType struct {
	Id                string `json:"id,omitempty"`
	IsProviderEnabled bool   `json:"is_provider_enabled,omitempty"`
	Title             string `json:"title,omitempty"`
}

type Operator struct {
	Comments         string `json:"comments,omitempty"`
	ContactEmail     string `json:"contact_email,omitempty"`
	FaultReportEmail string `json:"fault_report_email,omitempty"`
	Id               string `json:"id,omitempty"`
	// Is this operator a private individual vs a company
	IsPrivateIndividual bool   `json:"is_private_individual,omitempty"`
	PhonePrimary        string `json:"phone_primary,omitempty"`
	PhoneSecondary      string `json:"phone_secondary,omitempty"`
	Title               string `json:"title,omitempty"`
	Website             string `json:"website,omitempty"`
}

type Poi struct {
	// The address
	Address *Address `json:"address,omitempty"`
	// The connections available at this charge point
	Connections []Connection `json:"connections,omitempty"`
	// The cost of charging
	Cost string `json:"cost,omitempty"`
	// The ID of the data provider
	DataProviderId string `json:"data_provider_id,omitempty"`
	// The ID of the charger
	Id string `json:"id,omitempty"`
	// The number of charging points
	NumPoints int64 `json:"num_points,string,omitempty"`
	// The operator
	Operator *Operator `json:"operator,omitempty"`
	// The ID of the operator of the charger
	OperatorId string `json:"operator_id,omitempty"`
	// The type of usage
	UsageType *UsageType `json:"usage_type,omitempty"`
	// The type of usage for this charger point (is it public, membership required, etc)
	UsageTypeId string `json:"usage_type_id,omitempty"`
}

type ReferenceDataRequest struct {
}

type ReferenceDataResponse struct {
	// The types of charger
	ChargerTypes *ChargerType `json:"charger_types,omitempty"`
	// The types of checkin status
	CheckinStatusTypes *CheckinStatusType `json:"checkin_status_types,omitempty"`
	// The types of connection
	ConnectionTypes *ConnectionType `json:"connection_types,omitempty"`
	// The countries
	Countries []Country `json:"countries,omitempty"`
	// The types of current
	CurrentTypes *CurrentType `json:"current_types,omitempty"`
	// The providers of the charger data
	DataProviders *DataProvider `json:"data_providers,omitempty"`
	// The companies operating the chargers
	Operators []Operator `json:"operators,omitempty"`
	// The status of the charger
	StatusTypes *StatusType `json:"status_types,omitempty"`
	// The status of a submission
	SubmissionStatusTypes *SubmissionStatusType `json:"submission_status_types,omitempty"`
	// The different types of usage
	UsageTypes *UsageType `json:"usage_types,omitempty"`
	// The types of user comment
	UserCommentTypes *UserCommentType `json:"user_comment_types,omitempty"`
}

type SearchRequest struct {
	// Bounding box to search within (top left and bottom right coordinates)
	Box *BoundingBox `json:"box,omitempty"`
	// IDs of the connection type
	ConnectionTypes string `json:"connection_types,omitempty"`
	// Country ID
	CountryId string `json:"country_id,omitempty"`
	// Search distance from point in metres, defaults to 5000m
	Distance int64 `json:"distance,string,omitempty"`
	// Supported charging levels
	Levels []string `json:"levels,omitempty"`
	// Coordinates from which to begin search
	Location *Coordinates `json:"location,omitempty"`
	// Maximum number of results to return, defaults to 100
	MaxResults int64 `json:"max_results,string,omitempty"`
	// Minimum power in KW. Note: data not available for many chargers
	MinPower int64 `json:"min_power,string,omitempty"`
	// IDs of the the EV charger operator
	Operators []string `json:"operators,omitempty"`
	// Usage of the charge point (is it public, membership required, etc)
	UsageTypes string `json:"usage_types,omitempty"`
}

type SearchResponse struct {
	Pois []Poi `json:"pois,omitempty"`
}

type StatusType struct {
	Id            string `json:"id,omitempty"`
	IsOperational bool   `json:"is_operational,omitempty"`
	Title         string `json:"title,omitempty"`
}

type SubmissionStatusType struct {
	Id     string `json:"id,omitempty"`
	IsLive bool   `json:"is_live,omitempty"`
	Title  string `json:"title,omitempty"`
}

type UsageType struct {
	Id                   string `json:"id,omitempty"`
	IsAccessKeyRequired  bool   `json:"is_access_key_required,omitempty"`
	IsMembershipRequired bool   `json:"is_membership_required,omitempty"`
	IsPayAtLocation      bool   `json:"is_pay_at_location,omitempty"`
	Title                string `json:"title,omitempty"`
}

type UserCommentType struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
