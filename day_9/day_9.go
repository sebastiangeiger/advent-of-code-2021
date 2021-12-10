package day_9

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
		common.PrintNotImplemented(7, problem)
	}
}

func problem1() {
	fmt.Printf("RiskLevel (test): %d\n", solveProblem1("day_9_test.input"))
}

func problem2() {
	fmt.Println("Day 9 - Problem 2")
}

func solveProblem1(path string) int {
	matrix := readMatrix(path)
	dx := len(matrix)
	dy := len(matrix[0])
	lowPoints := []int{}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			current := matrix[x][y]
			adjacents := []int{}
			if y-1 >= 0 {
				//left
				adjacents = append(adjacents, matrix[x][y-1])
			}
			if y+1 < dy {
				//right
				adjacents = append(adjacents, matrix[x][y+1])
			}
			if x-1 >= 0 {
				//above
				adjacents = append(adjacents, matrix[x-1][y])
			}
			if x+1 < dx {
				//below
				adjacents = append(adjacents, matrix[x+1][y])
			}
			isLow := true
			for _, adjacent := range adjacents {
				if current >= adjacent {
					isLow = false
					break
				}
			}
			if isLow {
				lowPoints = append(lowPoints, current)
			}
		}
	}
	sum := 0
	for _, lowPoint := range lowPoints {
		sum += (lowPoint + 1)
	}

	return sum
}

func readMatrix(path string) [][]int {
	lines := common.ReadLinesFrom(path, false)
	matrix := [][]int{}
	for _, line := range lines {
		matrix = append(matrix, common.ToIntLine(line, ""))
	}
	return matrix
}
