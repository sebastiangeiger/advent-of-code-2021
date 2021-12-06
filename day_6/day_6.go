package day_6

import (
	"fmt"
	"strconv"
	"strings"

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
	population := common.ToIntLine(line, ",")
	maxDays := 18
	for day := 0; day <= maxDays; day++ {
		if printDebug {
			dayDisplay(day, population)
		}
		population = simulateDay(population)
	}
	return 1
}

func simulateDay(population []int) []int {
	return population
}

func dayDisplay(day int, population []int) {
	result := []string{}
	for _, pop := range population {
		result = append(result, strconv.Itoa(pop))
	}
	strResult := strings.Join(result, ",")
	fmt.Printf("After %2d days: %s\n", day, strResult)
}
