package processors

import (
	"example/hello/internal/models"
	"math/rand"
	"strings"
)

type PlanetProcessorInterface interface {
	FindMatchingPlanets(planets []models.Planet, weatherDescription string) []models.Planet
	SanitizeIncomingWords(input string) []string
	ClimateMatch(planetClimates, keywords []string) bool
	ConvertWeather(weather string, temp float64) string
}

type PlanetProcessor struct {}

func NewPlanetProcessor() PlanetProcessor {
	return PlanetProcessor{}
}

func (p *PlanetProcessor) FindMatchingPlanets(planets []models.Planet, weatherDescription string) []models.Planet {
	var matches []models.Planet
	keywords := p.SanitizeIncomingWords(weatherDescription)

	for _, planet := range planets {
		planetClimates := p.SanitizeIncomingWords(planet.Climate)
		if p.ClimateMatch(planetClimates, keywords) {
			matches = append(matches, planet)
		}
	}

	randomIndex := rand.Intn(len(matches))
	return matches[randomIndex:randomIndex+1]
}

func (p *PlanetProcessor) SanitizeIncomingWords(input string) []string {
	words := strings.Split(input, ",")
	list := make([]string, 0, len(words))

	for _, word := range words {
		sanitized := strings.TrimSpace(strings.ToLower(word))
		if sanitized != "" {
			list = append(list, sanitized)
		}
	}

	return list
}

func (p *PlanetProcessor) ClimateMatch(planetClimates, keywords []string) bool {
	for _, keyword := range keywords {
		for _, planetClimate := range planetClimates {
			if strings.Contains(planetClimate, keyword) {
				return true
			}
		}
	}

	return false
}

func (p *PlanetProcessor) ConvertWeather(weather string, temp float64) string {
	// First check extreme temperatures
	if temp < -30 {
		return "frozen"
	}
	if temp > 40 {
		return "superheated"
	}

	// Then check weather conditions with temperature ranges
	switch weather {
	case "snow", "blizzard":
		if temp > -5 {
		    return "subarctic"
		}
		if temp > -15 {
			return "frigid"
		}
		return "frozen"

	case "rain", "storm", "thunder":
		if temp > 25 {
			return "tropical"
		}
		return "temperate"

	case "clear", "sun":
		if temp > 30 {
			return "arid"
		}
		if temp < 0 {
			return "subarctic"
		}
		return "temperate"

	case "cloudy":
		if temp > 30 {
			return "arid"
		}
		if temp < 0 {
			return "frigid"
		}
		return "temperate"

	case "mist", "fog":
		if temp > 20 {
			return "humid"
		}
		if temp < 0 {
			return "frigid"
		}
		return "temperate"

	case "ash", "dust", "sand", "smoke":
		return "ash"

	default:
		// Temperature-only fallbacks
		if temp > 35 {
			return "arid"
		}
		if temp > 25 {
			return "tropical"
		}
		if temp < 0 {
			return "frigid"
		}
		return "temperate"
	}
}
