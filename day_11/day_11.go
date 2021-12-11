package day_11

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
		common.PrintNotImplemented(11, problem)
	}
}

func problem1() {
	fmt.Printf("Problem 1: %d", solveProblem1("day_11_test.input"))
}

func problem2() {
	fmt.Println("Day 11 - Problem 2")
}

func solveProblem1(path string) int {
	population := read(path)
	step(population)
	return 1
}

func read(path string) [][]int {
	result := [][]int{}
	for _, line := range common.ReadLinesFrom(path, false) {
		result = append(result, common.ToIntLine(line, ""))
	}
	return result
}

func step(population [][]int) {
	newPopulation := copyPopulation(population)
	increment(newPopulation)
	print(population)
	print(newPopulation)
}

func print(population [][]int) {
	for _, row := range population {
		for _, individual := range row {
			fmt.Printf("%d", individual)
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}

func increment(population [][]int) {
	for i, row := range population {
		for j := range row {
			population[i][j] += 1
		}
	}
}

func copyPopulation(population [][]int) [][]int {
	newPopulation := common.InitializeArray(len(population), len(population[0]))
	for i, row := range population {
		for j := range row {
			newPopulation[i][j] += population[i][j]
		}
	}
	return newPopulation
}
