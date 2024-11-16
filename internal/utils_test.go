package internal

import (
	"testing"
)

func TestConvertStringToFloat(t *testing.T) {
  tests := []struct {
		input    string
		expected float64
	}{
		{"123.456", 123.456},
		{"0", 0},
		{"-987.65", -987.65},
		{"123 TEST 123", 123},
		{"1.23e4", 1.23}, // Scientific notation
		{"invalid", 0},     // Invalid input
	}

  for _, test := range tests {
    result := ConvertStringToFloat(test.input)
		if result != test.expected {
			t.Errorf("ConvertStringToFloat(%q) = %v; want %v", test.input, result, test.expected)
		}
  }
}


func TestHasMin(t *testing.T) {
  tests := []struct {
		input    string
		expected bool
	}{
		{"minimum", true},
		{"work", false},
		{"muni", false},
		{"intermini", true},
	}

  for _, test := range tests {
    result := HasMin(test.input)
		if result != test.expected {
			t.Errorf("HasMin(%q) = %v; want %v", test.input, result, test.expected)
		}
  }
}
