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
	fmt.Printf("After 80 days (test): %d\n", solveProblem("day_6_test.input", 80, false))
	fmt.Printf("After 80 days (real): %d\n", solveProblem("day_6.input", 80, false))
}

func problem2() {
	fmt.Printf("After 256 days (test): %d\n", solveProblem("day_6_test.input", 256, true))
}

func solveProblem(path string, maxDays int, printDebug bool) int {
	line := common.ReadLinesFrom(path, false)[0]
	population := common.ToIntLine(line, ",")
	for day := 0; day < maxDays; day++ {
		if printDebug && day <= 18 {
			dayDisplay(day, population)
		} else if printDebug {
			fmt.Printf("Day %d\n", day)
		}
		population = simulateDay(population)
	}
	return len(population)
}

func simulateDay(population []int) []int {
	agedPopulation := make([]int, len(population))
	newSpawnsAmount := 0
	for i, individual := range population {
		if individual == 0 {
			agedPopulation[i] = 6
			newSpawnsAmount += 1
		} else {
			agedPopulation[i] = individual - 1
		}
	}
	newSpawns := make([]int, newSpawnsAmount)
	for i := 0; i < newSpawnsAmount; i++ {
		newSpawns[i] = 8
	}
	return append(agedPopulation, newSpawns...)
}

func dayDisplay(day int, population []int) {
	result := []string{}
	for _, pop := range population {
		result = append(result, strconv.Itoa(pop))
	}
	strResult := strings.Join(result, ",")
	fmt.Printf("After %2d days: %s\n", day, strResult)
}
