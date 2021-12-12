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
	fmt.Printf("Number of flashes (test): %d\n", solveProblem1("day_11_test.input"))
	fmt.Printf("Number of flashes (real): %d\n", solveProblem1("day_11.input"))
}

func problem2() {
	fmt.Printf("All flash in step (test): %d\n", solveProblem2("day_11_test.input"))
	fmt.Printf("All flash in step (real): %d\n", solveProblem2("day_11.input"))
}

func solveProblem1(path string) int {
	population := read(path)
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += step(population)
	}
	return flashes
}

func solveProblem2(path string) int {
	population := read(path)
	i := 0
	for {
		i++
		step(population)
		if allFlash(population) {
			break
		}
	}
	return i
}

func allFlash(population [][]int) bool {
	sum := 0
	for _, row := range population {
		for _, individual := range row {
			sum += individual
		}
	}
	return sum == 0
}

func read(path string) [][]int {
	result := [][]int{}
	for _, line := range common.ReadLinesFrom(path, false) {
		result = append(result, common.ToIntLine(line, ""))
	}
	return result
}

func step(population [][]int) int {
	increment(population)
	return flashes(population)
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

func flashes(population [][]int) int {
	result := 0
	for {
		incremented := false
		for x, row := range population {
			for y, energyLevel := range row {
				if energyLevel > 9 {
					result++
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
	return result
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
