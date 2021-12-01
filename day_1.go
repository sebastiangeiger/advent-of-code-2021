package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		printHelp()
	} else {
		day, dayErr := strconv.Atoi(os.Args[1])
		problem, problemErr := strconv.Atoi(os.Args[2])
		if dayErr != nil {
			fmt.Println("Day was not an integer", dayErr)
			os.Exit(1)
		} else if problemErr != nil {
			fmt.Println("Problem was not an integer", dayErr)
			os.Exit(1)
		} else {
			runDayProblem(day, problem)
		}
	}
}

func printHelp() {
	fmt.Printf("Usage: '%s day problem'", os.Args[0])
	os.Exit(1)
}

func printNotImplemented(day int, problem int) {
	fmt.Printf("Day %d - problem %d is not implemented yet", day, problem)
	os.Exit(1)
}

func runDayProblem(day int, problem int) {
	switch day {
	case 1:
		runDay1(problem)
	}
}

func runDay1(problem int) {
	switch problem {
	case 1:
		problem1()
	case 2:
		problem2()
	default:
		printNotImplemented(1, problem)
	}
}

func problem1() {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Printf("Increases (test): %d\n", countStepIncreases(testInput))
	input := readIntsFrom("day_1.input")
	fmt.Printf("Increases (real): %d\n", countStepIncreases(input))
}

func problem2() {
	fmt.Println("Please implement!")
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
