package day_3

import (
	"fmt"
	"math"
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
	pivoted := pivot(toIntMatrix(lines))
	gamma := calculateGamma(pivoted)
	fmt.Printf("pivoted: %#v\n", pivoted)
	fmt.Printf("gamma: %#v\n", makeBase10(gamma))
}

func calculateGamma(input [][]int) []int {
	output := make([]int, len(input))
	for i, numbers := range input {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		majority := 0
		if sum > len(numbers)/2.0 {
			majority = 1
		}
		output[i] = majority
	}
	return output
}

func makeBase10(input []int) int {
	length := len(input)
	output := 0
	for i, number := range input {
		significance := float64((length - 1) - i)
		output += int(math.Pow(2, significance) * float64(number))
	}
	return output
}

// This enforces that we get an array of arrays with the same dimensions and the values being 0 or 1
func toIntMatrix(lines []string) [][]int {
	dx := len(lines)
	dy := len(lines[0])
	result := initializeArray(dx, dy)
	for x := 0; x < dx; x++ {
		currentLine := lines[x]
		if len(currentLine) != dy {
			panic(fmt.Sprintf("Expected lines[%d] to be %d long but was %d", x, dy, len(currentLine)))
		}
		for y := 0; y < dy; y++ {
			number, err := strconv.Atoi(string(currentLine[y]))
			if err != nil {
				panic(err)
			} else if number != 0 && number != 1 {
				panic(fmt.Sprintf("Got a number that's not 0 or 1: %d", number))
			} else {
				result[x][y] = number
			}
		}
	}
	return result
}

func pivot(input [][]int) [][]int {
	dx := len(input[0])
	dy := len(input)
	output := initializeArray(dx, dy)
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			output[x][y] = input[y][x]
		}
	}
	return output
}

func initializeArray(dx int, dy int) [][]int {
	result := make([][]int, dx)
	for i := 0; i < dx; i++ {
		result[i] = make([]int, dy)
	}
	return result
}
