package prayer

import (
	"go.m3o.com/client"
)

type Prayer interface {
	Times(*TimesRequest) (*TimesResponse, error)
}

func NewPrayerService(token string) *PrayerService {
	return &PrayerService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type PrayerService struct {
	client *client.Client
}

// Get the prayer (salah) times for a location on a given date
func (t *PrayerService) Times(request *TimesRequest) (*TimesResponse, error) {

	rsp := &TimesResponse{}
	return rsp, t.client.Call("prayer", "Times", request, rsp)

}

type PrayerTime struct {
	// asr time
	Asr string `json:"asr,omitempty"`
	// date for prayer times in YYYY-MM-DD format
	Date string `json:"date,omitempty"`
	// fajr time
	Fajr string `json:"fajr,omitempty"`
	// isha time
	Isha string `json:"isha,omitempty"`
	// maghrib time
	Maghrib string `json:"maghrib,omitempty"`
	// time of sunrise
	Sunrise string `json:"sunrise,omitempty"`
	// zuhr time
	Zuhr string `json:"zuhr,omitempty"`
}

type TimesRequest struct {
	// optional date in YYYY-MM-DD format, otherwise uses today
	Date string `json:"date,omitempty"`
	// number of days to request times for
	Days int32 `json:"days,omitempty"`
	// optional latitude used in place of location
	Latitude float64 `json:"latitude,omitempty"`
	// location to retrieve prayer times for.
	// this can be a specific address, city, etc
	Location string `json:"location,omitempty"`
	// optional longitude used in place of location
	Longitude float64 `json:"longitude,omitempty"`
}

type TimesResponse struct {
	// date of request
	Date string `json:"date,omitempty"`
	// number of days
	Days int32 `json:"days,omitempty"`
	// latitude of location
	Latitude float64 `json:"latitude,omitempty"`
	// location for the request
	Location string `json:"location,omitempty"`
	// longitude of location
	Longitude float64 `json:"longitude,omitempty"`
	// prayer times for the given location
	Times []PrayerTime `json:"times,omitempty"`
}
