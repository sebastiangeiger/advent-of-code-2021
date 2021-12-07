package day_7

import (
	"fmt"
	"math"

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
	solution := solveProblem1("day_7_test.input")
	fmt.Printf("Minimum Fuel Cost (test): %d", solution)
}

func problem2() {
	fmt.Println("Day 7 - Problem 2")
}

func solveProblem1(path string) int {
	line := common.ReadLinesFrom(path, false)[0]
	positions := common.ToIntLine(line, ",")
	fuelCosts := []int{}
	for i := common.Min(positions...); i <= common.Max(positions...); i++ {
		fuelCosts = append(fuelCosts, alignTo(positions, i))
	}
	return common.Min(fuelCosts...)
}

func alignTo(positions []int, alignment int) int {
	cost := 0.0
	for _, position := range positions {
		cost += math.Abs(float64(position) - float64(alignment))
	}
	return int(cost)
}
