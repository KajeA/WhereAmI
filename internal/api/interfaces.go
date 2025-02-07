package api

import "example/hello/internal/models"

type WeatherAPI interface {
	GetWeather(city string) (*WeatherResponse, error)
}

type StarwarsAPI interface {
	GetPlanet() ([]models.Planet, error)
}
