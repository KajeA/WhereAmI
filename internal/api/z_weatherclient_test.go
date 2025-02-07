package api

import (
"testing"
)

func TestWeatherClient(t *testing.T) {
	client := NewWeatherClient("5e34fbe51acc18c4fcce6b3895aa05c3")

	city := "London"
	response, err := client.GetWeather(city)

	if err != nil {
		t.Fatalf("Failed to get weather: %v", err)
	}

	CheckWeatherResponse(t, response)
}

func CheckWeatherResponse(t *testing.T, response *WeatherResponse) {

	if response == nil {
		t.Fatal("No response from API")
	}

	if response.Name == "" {
		t.Error("No city input")
	}

	if len(response.Weather) == 0 {
		t.Error("No data retrieved")
	}

	if len(response.Weather) > 0 {
		if response.Weather[0].Description == "" {
			t.Error("No weather description")
		}
	}
}
