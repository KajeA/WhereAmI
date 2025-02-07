package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WeatherHTTPClient interface {
	GetWeather(url string) (*WeatherResponse, error)
}

type WeatherClient struct {
	APIKey string
}

func NewWeatherClient(apiKey string) WeatherHTTPClient {
	return &WeatherClient{APIKey: apiKey}
}

type WeatherResponse struct
{
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func (c *WeatherClient) GetWeather(city string) (*WeatherResponse, error) {
    encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", encodedCity, c.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	fmt.Printf("Response Body: %s\n", string(body))

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v", err)
	}

	if len(weather.Weather) == 0 {
		return nil, fmt.Errorf("no weather data in response")
	}

	return &weather, nil
}
