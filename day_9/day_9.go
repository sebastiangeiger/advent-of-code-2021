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

type Point struct {
	x int
	y int
}

func problem1() {
	fmt.Printf("RiskLevel (test): %d\n", solveProblem1("day_9_test.input"))
	fmt.Printf("RiskLevel (real): %d\n", solveProblem1("day_9.input"))
}

func problem2() {
	fmt.Printf("Basins (test): %d\n", solveProblem2("day_9_test.input"))
}

func findLowPoints(matrix [][]int) []Point {
	dx := len(matrix)
	dy := len(matrix[0])
	lowPoints := []Point{}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			current := Point{x, y}
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
				if matrix[current.x][current.y] >= adjacent {
					isLow = false
					break
				}
			}
			if isLow {
				lowPoints = append(lowPoints, current)
			}
		}
	}
	return lowPoints
}

func solveProblem1(path string) int {
	matrix := readMatrix(path)
	sum := 0
	for _, p := range findLowPoints(matrix) {
		sum += (matrix[p.x][p.y] + 1)
	}

	return sum
}

func solveProblem2(path string) int {
	matrix := readMatrix(path)
	basins := initializeBasins(matrix)
	fmt.Println(basins)
	return 1
}

func initializeBasins(matrix [][]int) [][]Point {
	basins := [][]Point{}
	for _, p := range findLowPoints(matrix) {
		basins = append(basins, []Point{p})
	}
	return basins
}

func readMatrix(path string) [][]int {
	lines := common.ReadLinesFrom(path, false)
	matrix := [][]int{}
	for _, line := range lines {
		matrix = append(matrix, common.ToIntLine(line, ""))
	}
	return matrix
}
