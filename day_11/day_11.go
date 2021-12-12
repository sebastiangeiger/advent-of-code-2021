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
	fmt.Println("Original Population")
	print(population)
	increment(population)
	flashes(population)
	fmt.Println("After Round 1")
	print(population)
	increment(population)
	flashes(population)
	fmt.Println("After Round 2")
	print(population)
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

func flashes(population [][]int) {
	for {
		incremented := false
		for x, row := range population {
			for y, energyLevel := range row {
				if energyLevel > 9 {
					incremented = true
					incrementNeighbors(population, x, y)
					population[x][y] = 0
				}
			}
		}
		if !incremented {
			break
		}
	}
}

func incrementNeighbors(population [][]int, x int, y int) {
	dx := len(population)
	dy := len(population[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			original := (i == 0 && j == 0)
			inBounds := x+i >= 0 && x+i < dx && y+j >= 0 && y+j < dy
			if !original && inBounds {
				currentValue := population[x+i][y+j]
				if currentValue > 0 && currentValue <= 10 {
					population[x+i][y+j] += 1
				}
			}
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
