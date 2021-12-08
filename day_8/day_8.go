package day_8

import (
	"fmt"
	"strings"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

type Observation struct {
	signals []string
	output  []string
}

func Run(problem int) {
	switch problem {
	case 1:
		problem1()
	case 2:
		problem2()
	default:
		common.PrintNotImplemented(8, problem)
	}
}

func interpretations(pattern string) []int {
	switch len(pattern) {
	case 2:
		return []int{1}
	case 3:
		return []int{7}
	case 4:
		return []int{4}
	case 5:
		return []int{2, 3, 5}
	case 6:
		return []int{0, 6, 9}
	case 7:
		return []int{8}
	}
	panic("Should not get here")
}

func problem1() {
	fmt.Printf("1,4,7,8 in output (test): %d", solveProblem1("day_8_test.input"))
}

func problem2() {
	fmt.Println("Day 8 - Problem 2")
}

func solveProblem1(path string) int {
	observations := parseLines(common.ReadLinesFrom(path, false))
	sum := 0
	for _, observation := range observations {
		for _, digit := range observation.output {
			interpretations := interpretations(digit)
			if len(interpretations) == 1 {
				sum += 1
			}
		}
	}
	return sum
}

func parseLines(lines []string) []Observation {
	observations := make([]Observation, len(lines))
	for i, line := range lines {
		modified := strings.Split(line, "|")
		if len(modified) == 2 {
			signalPatterns := parsePatterns(modified[0])
			output := parsePatterns(modified[1])
			observations[i] = Observation{signalPatterns, output}
		} else {
			panic("Expected 2")
		}
	}
	return observations
}

func parsePatterns(rawDigits string) []string {
	output := []string{}
	for _, digit := range strings.Split(rawDigits, " ") {
		if len(digit) > 0 {
			output = append(output, digit)
		}
	}
	return output
}
