package main

import "fmt"

func main() {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Printf("Increases: %d\n", countStepIncreases(testInput))
}

func countStepIncreases(measurements []int) int {
	previousMeasurement := measurements[0]
	increases := 0
	for _, measurement := range measurements[1:] {
		if previousMeasurement < measurement {
			increases += 1
		}
		previousMeasurement = measurement
	}
	return increases
}
