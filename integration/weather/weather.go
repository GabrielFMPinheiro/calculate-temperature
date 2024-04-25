package weather

import (
	"encoding/json"
	"net/http"
	"os"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Latitude       float64 `json:"lat"`
	Longitude      float64 `json:"lon"`
	TimeZoneID     string  `json:"tz_id"`
	LocalTimeEpoch int64   `json:"localtime_epoch"`
	LocalTime      string  `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

type CurrentWeather struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TemperatureC     float64   `json:"temp_c"`
	TemperatureF     float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMPH          float64   `json:"wind_mph"`
	WindKPH          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDirection    string    `json:"wind_dir"`
	PressureMB       float64   `json:"pressure_mb"`
	PressureIN       float64   `json:"pressure_in"`
	PrecipMM         float64   `json:"precip_mm"`
	PrecipIN         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	VisibilityKM     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UVIndex          float64   `json:"uv"`
	GustMPH          float64   `json:"gust_mph"`
	GustKPH          float64   `json:"gust_kph"`
}

type WeatherData struct {
	Location Location       `json:"location"`
	Current  CurrentWeather `json:"current"`
}

type Weather struct {
	Temp_C float64 `json:"temp_C"`
	Temp_F float64 `json:"temp_F"`
	Temp_K float64 `json:"temp_K"`
}

type WeatherIntegration struct{}

func (w *WeatherIntegration) GetWeather(city string) (*Weather, error) {
	// Here we would make a request to a weather API
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	city, _, _ = transform.String(t, city)

	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json?key="+os.Getenv("WEATHER_API_KEY")+"&q="+city, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Here we would parse the response from the weather API
	var weather WeatherData
	err = json.NewDecoder(resp.Body).Decode(&weather)

	if err != nil {
		return nil, err
	}

	// and return the temperature in Celsius, Fahrenheit and Kelvin
	return &Weather{
		Temp_C: weather.Current.TemperatureC,
		Temp_F: weather.Current.TemperatureF,
		Temp_K: weather.Current.TemperatureC + 273,
	}, nil
}
