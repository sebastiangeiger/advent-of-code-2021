package main

import (
	"fmt"
	"strconv"
)

func runDay1(problem int) {
	switch problem {
	case 1:
		day1problem1()
	case 2:
		day1problem2()
	default:
		printNotImplemented(1, problem)
	}
}

func day1problem1() {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Printf("Increases (test): %d\n", countStepIncreases(testInput))
	input := readIntsFrom("day_1.input")
	fmt.Printf("Increases (real): %d\n", countStepIncreases(input))
}

func day1problem2() {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Printf("Increases (test): %d\n", countStepIncreases(slidingAverages(testInput, 3)))
	input := readIntsFrom("day_1.input")
	fmt.Printf("Increases (real): %d\n", countStepIncreases(slidingAverages(input, 3)))
}

func slidingAverages(measurements []int, window int) []int {
	result := []int{}
	for i := 0; i <= len(measurements)-window; i++ {
		average := 0
		for j := 0; j < window; j++ {
			average += measurements[i+j]
		}
		result = append(result, average)
	}
	return result
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
	lines := readLinesFrom(path)
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
