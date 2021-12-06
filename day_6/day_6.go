package day_6

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	fmt.Printf("After 80 days (test): %d\n", solveProblem1("day_6_test.input", 80, false))
	fmt.Printf("After 80 days (real): %d\n", solveProblem1("day_6.input", 80, false))
}

func problem2() {
	fmt.Printf("After 256 days (test): %d\n", solveProblem2("day_6_test.input", 256, true))
}

func solveProblem1(path string, maxDays int, printDebug bool) int {
	line := common.ReadLinesFrom(path, false)[0]
	population := toInt8Array(common.ToIntLine(line, ","))
	table := buildLookupTable(80)
	sum := 0
	for _, individual := range population {
		sum += table[individual][80]
	}
	return sum
}

func solveProblem2(path string, maxDays int, printDebug bool) int {
	line := common.ReadLinesFrom(path, false)[0]
	population := toInt8Array(common.ToIntLine(line, ","))
	fmt.Printf("%#v\n", population)
	table := buildLookupTable(80)
	fmt.Printf("   ")
	for day, _ := range table[0] {
		fmt.Printf("%4d ", day)
	}
	fmt.Printf("\n")
	for num, row := range table {
		fmt.Printf("%d: ", num)
		for _, cell := range row {
			fmt.Printf("%4d ", cell)
		}
		fmt.Printf("\n")
	}
	return 1
}

func buildLookupTable(maxDays int) [][]int {
	result := common.InitializeArray(9, maxDays+1)
	for num := 0; num <= 8; num++ {
		for day := 0; day <= maxDays; day++ {
			res := simulatePopulation([]int8{int8(num)}, day, false)
			result[num][day] = res
		}
	}
	return result
}

func simulatePopulation(population []int8, maxDays int, printDebug bool) int {
	for day := 0; day < maxDays; day++ {
		start := time.Now()
		if printDebug && day <= 18 {
			dayDisplay(day, population)
		} else if printDebug {
			fmt.Printf("Day %d: %d", day, len(population))
		}
		population = simulateDay(population)
		elapsed := time.Since(start)
		if printDebug {
			fmt.Printf(" took %s\n", elapsed)
		}
	}
	return len(population)
}

func simulateDay(population []int8) []int8 {
	newSpawnsAmount := 0
	for i, individual := range population {
		if individual == 0 {
			population[i] = 6
			newSpawnsAmount += 1
		} else {
			population[i] = individual - 1
		}
	}
	newSpawns := make([]int8, newSpawnsAmount)
	for i := 0; i < newSpawnsAmount; i++ {
		newSpawns[i] = 8
	}
	return append(population, newSpawns...)
}

func toInt8Array(array []int) []int8 {
	result := make([]int8, len(array))
	for i, el := range array {
		result[i] = int8(el)
	}
	return result
}

func dayDisplay(day int, population []int8) {
	result := []string{}
	for _, pop := range population {
		result = append(result, strconv.Itoa(int(pop)))
	}
	strResult := strings.Join(result, ",")
	fmt.Printf("After %2d days: %s\n", day, strResult)
}
