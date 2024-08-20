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
    var xValues []float64

    for scanner.Scan() {
        line := scanner.Text()

        value, err := strconv.ParseFloat(line, 64)
        if err != nil {
            fmt.Printf("Error parsing input: %v\n", err)
            continue
        }
        data = append(data, value)
        xValues = append(xValues, float64(len(data)))

        if len(data) > 1 {
            beta1, beta0 := stats.CalculateLinearRegression(xValues, data)

            nextX := float64(len(data) + 1)
            prediction := beta1*nextX + beta0

            stDev := math.Sqrt(stats.Variance(data))

            zScore := 1.96 // for 95% confidence interval
            lowerBound := prediction - zScore*stDev
            upperBound := prediction + zScore*stDev

            fmt.Printf("%.2f %.2f\n", lowerBound, upperBound)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error scanning input\n")
        return
    }
}
