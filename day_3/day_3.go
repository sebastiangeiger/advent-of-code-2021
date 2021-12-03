package day_3

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
		common.PrintNotImplemented(2, problem)
	}
}

func problem1() {
	solveProblem1("day_3_test.input")
}

func problem2() {
	fmt.Printf("Implement me!\n")
}

func solveProblem1(path string) {
	lines := common.ReadLinesFrom(path)
	pivoted := toIntArrays(lines)
	fmt.Printf("%#v\n", pivoted)
}

func toIntArrays(lines []string) [][]int {
	dx := len(lines)
	dy := len(lines[0])
	result := initializeArray(dx, dy)
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			number, err := strconv.Atoi(string(lines[x][y]))
			if err == nil {
				result[x][y] = number
			} else {
				panic(err)
			}
		}
	}
	return result
}

func initializeArray(dx int, dy int) [][]int {
	result := make([][]int, dx)
	for i := 0; i < dx; i++ {
		result[i] = make([]int, dy)
	}
	return result
}
