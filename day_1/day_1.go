package day_1

import (
	"fmt"
	"strconv"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

func Run(problem int) {
	switch problem {
	case 1:
		problem1()
	case 2:
		problem2()
	default:
		common.PrintNotImplemented(1, problem)
	}
}

func problem1() {
	testInput := readIntsFrom("day_1_test.input")
	fmt.Printf("Increases (test): %d\n", countStepIncreases(testInput))
	input := readIntsFrom("day_1.input")
	fmt.Printf("Increases (real): %d\n", countStepIncreases(input))
}

func problem2() {
	testInput := readIntsFrom("day_1_test.input")
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
	lines := common.ReadLinesFrom(path, false)
	integers := []int{}
	for _, line := range lines {
		integer, err := strconv.Atoi(line)
		if err == nil {
			integers = append(integers, integer)
		} else {
			panic(err)
		}
	}
	return integers
}
