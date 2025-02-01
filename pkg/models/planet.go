package models

type Planet struct {
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

type PlanetResponse struct {
	Next    string   `json:"next"`
	Results []Planet `json:"results"`
}
