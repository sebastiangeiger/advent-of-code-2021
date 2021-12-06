package day_6

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
		common.PrintNotImplemented(6, problem)
	}
}

func problem1() {
	fmt.Printf("After 80 days (test): %d", solveProblem1("day_6_test.input", true))
}

func problem2() {
	fmt.Println("Day 6 - Problem 2")
}

func solveProblem1(path string, printDebug bool) int {
	line := common.ReadLinesFrom(path, false)[0]
	initialPop := common.ToIntLine(line, ",")
	if printDebug {
		fmt.Printf("%#v\n", initialPop)
	}
	return 1
}
