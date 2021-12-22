package day_22

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
		common.PrintNotImplemented(22, problem)
	}
}

func problem1() {
	fmt.Println("Day 22 - Problem 1")
}

func problem2() {
	fmt.Println("Day 22 - Problem 2")
}