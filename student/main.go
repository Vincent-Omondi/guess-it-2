package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/Vincent-Omondi/guess-it-2/stats"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data []float64
	windowSize := 5

	for scanner.Scan() {
		line := scanner.Text()

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}
		data = append(data, value)

		if len(data) > windowSize {
			data = data[len(data)-windowSize:]
		}

		if len(data) > 1 {
			mean := stats.Mean(data)
			stDev := math.Sqrt(stats.Variance(data))
			slope := calculateSlope(data)

			// Dynamic adjustment
			adjustment := slope * float64(windowSize-1)
			lowerBound := int(mean + adjustment - 1.5*stDev)
			upperBound := int(mean + adjustment + 1.5*stDev)

			fmt.Printf("%d %d\n", lowerBound, upperBound)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning input\n")
		return
	}
}

// Calculate the slope of the best-fit line for the data
func calculateSlope(data []float64) float64 {
	n := float64(len(data))
	if n < 2 {
		return 0.0
	}

	sumX, sumY, sumXY, sumX2 := 0.0, 0.0, 0.0, 0.0
	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	numerator := n*sumXY - sumX*sumY
	denominator := n*sumX2 - sumX*sumX

	if denominator == 0 {
		return 0.0
	}

	return numerator / denominator
}
