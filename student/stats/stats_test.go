package stats

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{"Positive numbers", []float64{1, 2, 3, 4, 5}, 3},
		{"Mixed numbers", []float64{-1, -2, 3, 4}, 1},
		{"Single value", []float64{5}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mean(tt.data)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{"Positive numbers", []float64{1, 2, 3, 4, 5}, 2.5},
		{"Mixed numbers", []float64{-1, -2, 3, 4}, 8.666666666666666},
		{"Single value", []float64{5}, 0}, // Variance for a single value should be 0
		{"Empty slice", []float64{}, 0},   // Handle empty case, expect 0 or error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Variance(tt.data)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCalculateLinearRegression(t *testing.T) {
	tests := []struct {
		name      string
		x         []float64
		y         []float64
		expectedB1 float64
		expectedB0 float64
	}{
		{"Simple linear relationship", []float64{1, 2, 3}, []float64{2, 4, 6}, 2, 0},
		{"Negative slope", []float64{1, 2, 3}, []float64{6, 4, 2}, -2, 8},
		{"No relationship", []float64{1, 2, 3}, []float64{3, 3, 3}, 0, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b1, b0 := CalculateLinearRegression(tt.x, tt.y)
			if math.Abs(b1-tt.expectedB1) > 1e-9 || math.Abs(b0-tt.expectedB0) > 1e-9 {
				t.Errorf("expected b1 = %v, b0 = %v; got b1 = %v, b0 = %v", tt.expectedB1, tt.expectedB0, b1, b0)
			}
		})
	}
}
