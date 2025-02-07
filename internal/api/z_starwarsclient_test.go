package api

import (
	"testing"
)

func TestSWClient(t *testing.T) {
	client := NewStarwarsClient(nil)
	planets, err := client.GetPlanet()

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(planets) == 0 {
		t.Error("Expected one  planet, got none")
	}

	found := false
	for _, p := range planets {
		if p.Name == "Tatooine" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected Tatooine")
	}
}
