package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Printf("Increases (test): %d\n", countStepIncreases(testInput))
	input := readIntsFrom("day_1.input")
	fmt.Printf("Increases (real): %d\n", countStepIncreases(input))
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

func readIntsFrom(path string) []int {
	data, fileError := os.ReadFile(path)
	if fileError != nil {
		panic(fileError)
	}
	lines := strings.Split(string(data), "\n")
	integers := []int{}
	for _, line := range lines {
		if line != "" {
			integer, err := strconv.Atoi(line)
			if err == nil {
				integers = append(integers, integer)
			} else {
				panic(err)
			}
		}
	}
	return integers
}
