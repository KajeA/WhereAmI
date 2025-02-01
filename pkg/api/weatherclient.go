package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type WeatherClient struct {
	APIKey string
}

func NewWeatherClient(apiKey string) *WeatherClient {
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

	var weather WeatherResponse

	return &weather, nil
}
