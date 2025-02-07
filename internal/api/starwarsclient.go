package api

import (
	"encoding/json"
	"example/hello/internal/models"
	"fmt"
	"io"
	"net/http"
)

type SwHTTPClient interface {
	Get(url string) (*http.Response, error)
}

type PlanetService interface {
	GetPlanet() ([]models.Planet, error)
}

type StarwarsClient struct {
	baseURL    string
	httpClient SwHTTPClient
}

func NewStarwarsClient(httpClient SwHTTPClient) PlanetService {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &StarwarsClient{
		baseURL:    "https://swapi.dev/api",
		httpClient: httpClient,
	}
}

func (c *StarwarsClient) GetPlanet() ([]models.Planet, error) {
	var allPlanets []models.Planet
	nextURL := fmt.Sprintf("%s/planets/", c.baseURL)

	for nextURL != "" {
		resp, err := c.httpClient.Get(nextURL)
		if err != nil {
			return nil, fmt.Errorf("error fetching planets: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response: %v", err)
		}

		var planetResp models.PlanetResponse
		if err := json.Unmarshal(body, &planetResp); err != nil {
			return nil, fmt.Errorf("error parsing JSON: %v", err)
		}

		allPlanets = append(allPlanets, planetResp.Results...)
		nextURL = planetResp.Next
	}

	return allPlanets, nil
}
