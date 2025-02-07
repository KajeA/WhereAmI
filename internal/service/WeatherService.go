package service

import (
	"example/hello/internal/api"
	"example/hello/internal/models"
	"example/hello/internal/processors"
	"fmt"
	"math"
)

type WeatherService struct {
	weatherClient api.WeatherAPI
	starwarsClient api.StarwarsAPI
	planetProcessor processors.PlanetProcessorInterface
}

func NewWeatherService(
	wc api.WeatherAPI,
	sc api.StarwarsAPI,
	pp processors.PlanetProcessorInterface,
) *WeatherService {
	return &WeatherService{
		weatherClient: wc,
		starwarsClient: sc,
		planetProcessor: pp,
	}
}

func (s *WeatherService) GetMatchingPlanet(city string) (models.Planet, error) {
	// Get weather for city
	weather, err := s.weatherClient.GetWeather(city)
	if err != nil {
		return models.Planet{}, err
	}

	// Get planets
	planets, err := s.starwarsClient.GetPlanet()
	if err != nil {
		return models.Planet{}, err
	}

	// Convert weather to climate
	climate := s.planetProcessor.ConvertWeather(
		weather.Weather[0].Description,
		math.Round(weather.Main.Temp - 273.15), //convert from Kelvin
	)

	// Find matching planet
	matchingPlanets := s.planetProcessor.FindMatchingPlanets(planets, climate)
	if len(matchingPlanets) == 0 {
		return models.Planet{}, fmt.Errorf("no matching planet found")
	}

	return matchingPlanets[0], nil
}
