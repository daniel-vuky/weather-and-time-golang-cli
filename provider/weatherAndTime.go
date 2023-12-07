package weatherAndTime

import (
	"encoding/json"
	"fmt"
	"github.com/daniel-vuky/weather-and-time-golang-cli/config"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type WeatherData struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TZID           string  `json:"tz_id"`
		LocaltimeEpoch int64   `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`

	Current struct {
		LastUpdatedEpoch int64 `json:"last_updated_epoch"`
		LastUpdated      string
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`

		Condition struct {
			Text string `json:"text"`
			Code int    `json:"code"`
		} `json:"condition"`

		Wind struct {
			Mph     float64 `json:"wind_mph"`
			Kph     float64 `json:"wind_kph"`
			Degree  int     `json:"wind_degree"`
			Dir     string  `json:"wind_dir"`
			GustMph float64 `json:"gust_mph"`
			GustKph float64 `json:"gust_kph"`
		} `json:"wind"`

		PressureMB float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMM   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKM      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		UV         float64 `json:"uv"`
	} `json:"current"`
}

func GetWeatherAndTime(cityName string) {
	uri := getUri()
	params := getParams(cityName)
	// Prepare api url and call
	apiUrl := fmt.Sprintf("%s?%s", uri, params.Encode())
	response, responseError := http.Get(apiUrl)
	panicTheError(responseError)

	// Read the response body
	defer response.Body.Close()
	body, bodyErr := io.ReadAll(response.Body)
	panicTheError(bodyErr)

	// Convert the body to struct type
	var bodyConverted WeatherData
	bodyConvertedError := json.Unmarshal(body, &bodyConverted)
	panicTheError(bodyConvertedError)

	// Print the response body
	fmt.Printf("Time and the weather of %s is loading... \n", bodyConverted.Location.Name)
	fmt.Printf("Time: %s\n", bodyConverted.Location.Localtime)
	fmt.Printf("Weather: \n")
	printTheWeatherDetails(bodyConverted, "-")
}

func panicTheError(error error) {
	if error != nil {
		panic(error)
	}
}

func getUri() string {
	return config.GetUri()
}

func getApiKey() string {
	return config.GetApiKey()
}

func getParams(cityName string) url.Values {
	params := url.Values{}
	params.Add("key", getApiKey())
	params.Add("q", cityName)
	return params
}

func printTheWeatherDetails(bodyConverted interface{}, prefix string) {
	currentWeather := reflect.ValueOf(bodyConverted)
	for i := 0; i < currentWeather.NumField(); i++ {
		field := currentWeather.Field(i)
		fieldName := currentWeather.Type().Field(i).Name

		// If the field is a struct, recursively iterate through its fields
		if field.Kind() == reflect.Struct {
			printTheWeatherDetails(field.Interface(), fmt.Sprintf("%s%s", prefix, prefix))
		} else {
			// Print the field name and value
			fmt.Printf("%s%s: %v\n", prefix, fieldName, field.Interface())
		}
	}
}
