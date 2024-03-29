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
	fmt.Printf("After 80 days (test): %d\n", solveProblem("day_6_test.input", 80, false))
	fmt.Printf("After 80 days (real): %d\n", solveProblem("day_6.input", 80, false))
}

func problem2() {
	fmt.Printf("After 256 days (test): %d\n", solveProblem("day_6_test.input", 256, true))
	fmt.Printf("After 256 days (real): %d\n", solveProblem("day_6.input", 256, false))
}

func solveProblem(path string, maxDays int, printDebug bool) int {
	line := common.ReadLinesFrom(path, false)[0]
	population := common.ToIntLine(line, ",")
	table := buildLookupTable(maxDays)
	sum := 0
	for _, individual := range population {
		sum += table[individual][maxDays]
	}
	return sum
}

func buildLookupTable(maxDays int) [][]int {
	result := common.InitializeArray(9, maxDays+1)
	for day := 0; day <= maxDays; day++ {
		for num := 0; num <= 8; num++ {
			if day == 0 {
				// fmt.Printf("result[%d][%d] = 1\n", num, day)
				result[num][day] = 1
			} else if num == 0 {
				// fmt.Printf("result[%d][%d] = result[6][%d] + result[8][%d]\n", num, day, day-1, day-1)
				result[num][day] = result[6][day-1] + result[8][day-1]
			} else {
				result[num][day] = result[num-1][day-1]
			}
		}
	}
	return result
}
