package day_4

import (
	"fmt"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

func Run(problem int) {
	switch problem {
	case 1:
		problem1()
	case 2:
		problem2()
	default:
		common.PrintNotImplemented(4, problem)
	}
}

func problem1() {
	solveProblem1("day_4_test.input")
}

func problem2() {
	fmt.Println("Implement Day 4 - Problem 2")
}

func solveProblem1(path string) {
	lines := common.ReadLinesFrom(path, true)
	partitions := makePartitions(lines)
	fmt.Printf("partitions: %#v", partitions)
}

func makePartitions(lines []string) [][]string {
	output := [][]string{}
	currentPartition := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			output = append(output, currentPartition)
			currentPartition = []string{}
		} else {
			currentPartition = append(currentPartition, line)
		}
	}
	if len(currentPartition) > 0 {
		output = append(output, currentPartition)
	}
	return output
}
