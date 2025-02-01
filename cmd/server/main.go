package main

import (
	"example/hello/pkg/api"
	"example/hello/pkg/models"
	"example/hello/pkg/processors"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"path/filepath"
)

var (
	weatherClient *api.WeatherClient
	starwarsClient *api.StarwarsClient
	planetProcessor *processors.PlanetProcessor
)

type PageData struct {
	City        string
	WeatherData *api.WeatherResponse
	Temperature float64
	Planets     []models.Planet
	Error       string
}

func GetWeather(city string) (*api.WeatherResponse, error) {
	return weatherClient.GetWeather(city)
}

func main() {
	apiKey := "5e34fbe51acc18c4fcce6b3895aa05c3"

	weatherClient = api.NewWeatherClient(apiKey)
	starwarsClient = api.NewStarwarsClient()
	planetProcessor = processors.NewPlanetProcessor()

	rootDir, _ := filepath.Abs(".")
	staticDir := filepath.Join(rootDir, "../..", "web/static")
	templatesDir := filepath.Join(rootDir, "../..", "web/templates")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			City: r.URL.Query().Get("city"),
		}

		if data.City != "" {
			weatherData, err := GetWeather(data.City)
			if err != nil {
				data.Error = fmt.Sprintf("Error fetching weather data: %v", err)
			} else {
				data.WeatherData = weatherData
				data.Temperature = math.Round(weatherData.Main.Temp - 273.15)

				climate := planetProcessor.ConvertWeather(weatherData.Weather[0].Description, data.Temperature)
				planets, err := starwarsClient.GetPlanet()
				if err != nil {
					data.Error = fmt.Sprintf("Error fetching planets: %v", err)
				} else {
					data.Planets = planetProcessor.FindMatchingPlanets(planets, climate)
				}
			}
		}

		tmpl, err := template.ParseFiles(filepath.Join(templatesDir, "index.html"))
		if err != nil {
			log.Printf("Template parsing error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Template execution error: %v", err)
			// Don't write another header here, just log the error
			log.Printf("Error rendering template: %v", err)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	////Weather API call
	//http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
	//	city := r.URL.Query().Get("city")
	//
	//	weatherData, err := weatherClient.GetWeather(city)
	//	if err != nil {
	//		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
	//		return
	//	}
	//	w.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(weatherData)
	//})
	//
	//
	////Starwars API call
	//http.HandleFunc("/starwars", func(w http.ResponseWriter, r *http.Request) {
	//	starwarsData, err := starwarsClient.GetPlanet()
	//	if err != nil {
	//		http.Error(w, "Error fetching starwars data", http.StatusInternalServerError)
	//		return
	//	}
	//	w.Header().Set("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(starwarsData)
	//})
}
