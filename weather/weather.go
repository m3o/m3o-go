package weather

import (
	"go.m3o.com/client"
)

type Weather interface {
	Forecast(*ForecastRequest) (*ForecastResponse, error)
	Now(*NowRequest) (*NowResponse, error)
}

func NewWeatherService(token string) *WeatherService {
	return &WeatherService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type WeatherService struct {
	client *client.Client
}

// Get the weather forecast for the next 1-10 days
func (t *WeatherService) Forecast(request *ForecastRequest) (*ForecastResponse, error) {

	rsp := &ForecastResponse{}
	return rsp, t.client.Call("weather", "Forecast", request, rsp)

}

// Get the current weather report for a location by postcode, city, zip code, ip address
func (t *WeatherService) Now(request *NowRequest) (*NowResponse, error) {

	rsp := &NowResponse{}
	return rsp, t.client.Call("weather", "Now", request, rsp)

}

type Forecast struct {
	// the average temp in celsius
	AvgTempC float64 `json:"avg_temp_c,omitempty"`
	// the average temp in fahrenheit
	AvgTempF float64 `json:"avg_temp_f,omitempty"`
	// chance of rain (percentage)
	ChanceOfRain int32 `json:"chance_of_rain,omitempty"`
	// forecast condition
	Condition string `json:"condition,omitempty"`
	// date of the forecast
	Date string `json:"date,omitempty"`
	// the URL of forecast condition icon. Simply prefix with either http or https to use it
	IconUrl string `json:"icon_url,omitempty"`
	// max temp in celsius
	MaxTempC float64 `json:"max_temp_c,omitempty"`
	// max temp in fahrenheit
	MaxTempF float64 `json:"max_temp_f,omitempty"`
	// max wind speed kph
	MaxWindKph float64 `json:"max_wind_kph,omitempty"`
	// max wind speed mph
	MaxWindMph float64 `json:"max_wind_mph,omitempty"`
	// minimum temp in celsius
	MinTempC float64 `json:"min_temp_c,omitempty"`
	// minimum temp in fahrenheit
	MinTempF float64 `json:"min_temp_f,omitempty"`
	// time of sunrise
	Sunrise string `json:"sunrise,omitempty"`
	// time of sunset
	Sunset string `json:"sunset,omitempty"`
	// will it rain
	WillItRain bool `json:"will_it_rain,omitempty"`
}

type ForecastRequest struct {
	// number of days. default 1, max 10
	Days int32 `json:"days,omitempty"`
	// location of the forecase
	Location string `json:"location,omitempty"`
}

type ForecastResponse struct {
	// country of the request
	Country string `json:"country,omitempty"`
	// forecast for the next number of days
	Forecast []Forecast `json:"forecast,omitempty"`
	// e.g 37.55
	Latitude float64 `json:"latitude,omitempty"`
	// the local time
	LocalTime string `json:"local_time,omitempty"`
	// location of the request
	Location string `json:"location,omitempty"`
	// e.g -77.46
	Longitude float64 `json:"longitude,omitempty"`
	// region related to the location
	Region string `json:"region,omitempty"`
	// timezone of the location
	Timezone string `json:"timezone,omitempty"`
}

type NowRequest struct {
	// location to get weather e.g postcode, city
	Location string `json:"location,omitempty"`
}

type NowResponse struct {
	// cloud cover percentage
	Cloud int32 `json:"cloud,omitempty"`
	// the weather condition
	Condition string `json:"condition,omitempty"`
	// country of the request
	Country string `json:"country,omitempty"`
	// whether its daytime
	Daytime bool `json:"daytime,omitempty"`
	// feels like in celsius
	FeelsLikeC float64 `json:"feels_like_c,omitempty"`
	// feels like in fahrenheit
	FeelsLikeF float64 `json:"feels_like_f,omitempty"`
	// the humidity percentage
	Humidity int32 `json:"humidity,omitempty"`
	// the URL of the related icon. Simply prefix with either http or https to use it
	IconUrl string `json:"icon_url,omitempty"`
	// e.g 37.55
	Latitude float64 `json:"latitude,omitempty"`
	// the local time
	LocalTime string `json:"local_time,omitempty"`
	// location of the request
	Location string `json:"location,omitempty"`
	// e.g -77.46
	Longitude float64 `json:"longitude,omitempty"`
	// region related to the location
	Region string `json:"region,omitempty"`
	// temperature in celsius
	TempC float64 `json:"temp_c,omitempty"`
	// temperature in fahrenheit
	TempF float64 `json:"temp_f,omitempty"`
	// timezone of the location
	Timezone string `json:"timezone,omitempty"`
	// wind degree
	WindDegree int32 `json:"wind_degree,omitempty"`
	// wind direction
	WindDirection string `json:"wind_direction,omitempty"`
	// wind in kph
	WindKph float64 `json:"wind_kph,omitempty"`
	// wind in mph
	WindMph float64 `json:"wind_mph,omitempty"`
}
