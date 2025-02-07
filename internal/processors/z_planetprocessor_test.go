package processors

import (
	"testing"
)

func TestPlanetProcessor_SanitizeIncomingWords(t *testing.T) {
	processor := NewPlanetProcessor()

	tests := []struct {
		name string
		input string
		expected []string
	}{
		{
			name:     "multiple words",
			input:    "warm,   cold, snow",
			expected: []string{"warm", "cold", "snow"},
		},
		{
			name:     "weird spaces",
			input:    "warm,   cold , snow",
			expected: []string{"warm", "cold", "snow"},
		},
		{
			name:     "random caps",
			input:    "WArm, COLD, Snow",
			expected: []string{"warm", "cold", "snow"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.SanitizeIncomingWords(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %v words, got %v", len(tt.expected), len(result))
			}
			for i, word := range result {
				if word != tt.expected[i] {
					t.Errorf("Expected %v, got %v", tt.expected[i], word)
				}
			}
		})
	}
}
