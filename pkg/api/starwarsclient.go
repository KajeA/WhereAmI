package api

import (
	"encoding/json"
	"example/hello/pkg/models"
	"fmt"
	"io"
	"net/http"
)

type StarwarsClient struct {
	baseURL string
}

func NewStarwarsClient() *StarwarsClient {
	return &StarwarsClient{
		baseURL: "https://swapi.dev/api",
	}
}

func (c *StarwarsClient) GetPlanet() ([]models.Planet, error) {
	var allPlanets []models.Planet
	nextURL := fmt.Sprintf("%s/planets/", c.baseURL)

	for nextURL != "" {
		resp, err := http.Get(nextURL)
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
