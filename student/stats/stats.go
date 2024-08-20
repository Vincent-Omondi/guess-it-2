package stats

import (
	"fmt"
	"math"
)

func Mean(data []float64) float64 {
    if len(data) == 0 {
        fmt.Println("Cannot calculate mean of an empty slice")
    }
    sum := 0.0
    for _, value := range data {
        sum += value
    }
    return sum / float64(len(data))
}

func Variance(data []float64) float64 {
    if len(data) < 2 {
        fmt.Println("Variance requires at least two data points")
    }
    mean := Mean(data)
    var variance float64
    for _, value := range data {
        variance += (value - mean) * (value - mean)
    }
    return variance / float64(len(data)-1) // Use Bessel's correction
}

func CalculateLinearRegression(x, y []float64) (float64, float64) {
    if len(x) != len(y) {
        fmt.Println("x and y slices must have the same length")
    }
    n := float64(len(x))
    var sumX, sumY, sumXY, sumX2 float64

    for i := 0; i < len(x); i++ {
        sumX += x[i]
        sumY += y[i]
        sumXY += x[i] * y[i]
        sumX2 += x[i] * x[i]
    }

    denominator := n*sumX2 - sumX*sumX
    epsilon := 1e-9
    if math.Abs(denominator) < epsilon {
        fmt.Println("Denominator is too small, possibly causing division by zero")
    }

    beta1 := (n*sumXY - sumX*sumY) / denominator
    beta0 := (sumY - beta1*sumX) / n

    return beta1, beta0
}