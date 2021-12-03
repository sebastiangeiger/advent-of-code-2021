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
	fmt.Printf("Power (test): %d\n", solveProblem1("day_3_test.input"))
	fmt.Printf("Power (real): %d\n", solveProblem1("day_3.input"))
}

func problem2() {
	solveProblem2("day_3_test.input")
}

func solveProblem1(path string) int {
	lines := common.ReadLinesFrom(path)
	pivoted := pivot(toIntMatrix(lines))
	gamma := calculateGamma(pivoted)
	epsilon := complement(gamma)
	fmt.Printf("gamma: %#v\n", makeBase10(gamma))
	fmt.Printf("epsilon: %#v\n", makeBase10(epsilon))
	return makeBase10(gamma) * makeBase10(epsilon)
}

func solveProblem2(path string) {
	lines := common.ReadLinesFrom(path)
	matrix := toIntMatrix(lines)
	oxygenGeneratorRating := calculateOxygenGeneratorRating(matrix)
	fmt.Printf("oxygenGeneratorRating: %#v\n", makeBase10(oxygenGeneratorRating))
}

func calculateOxygenGeneratorRating(candidates [][]int) []int {
	for bitCriteriaIndex := 0; bitCriteriaIndex < len(candidates[0]); bitCriteriaIndex++ {
		pivoted := pivot(candidates)
		mostCommonValue := mostCommonValue(pivoted[bitCriteriaIndex])
		newCandidates := [][]int{}
		for _, candidate := range candidates {
			if candidate[bitCriteriaIndex] == mostCommonValue {
				newCandidates = append(newCandidates, candidate)
			}
		}
		// fmt.Printf("Most common value for '%#v' is '%#v'\n", pivoted[bitCriteriaIndex], mostCommonValue)
		// fmt.Printf("  %#v\ngot reduced to \n  %#v\n", candidates, newCandidates)
		if len(newCandidates) == 1 {
			return newCandidates[0]
		} else {
			candidates = newCandidates
		}
	}
	panic("I shouldn't get here")
}

func sum(input []int) int {
	sum := 0
	for _, number := range input {
		sum += number
	}
	return sum
}

func mostCommonValue(input []int) int {
	sum := sum(input)
	isEvenNumberItems := (len(input)%2 == 0)
	if isEvenNumberItems {
		if sum >= len(input)/2 {
			return 1
		} else {
			return 0
		}
	} else {
		// odd -> 7/2 = 3 meaning you need 4 items to be 1 for 1 to be most common
		if sum >= (len(input)/2)+1 {
			return 1
		} else {
			return 0
		}
	}
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

func complement(input []int) []int {
	output := make([]int, len(input))
	for i, number := range input {
		if number == 0 {
			output[i] = 1
		} else if number == 1 {
			output[i] = 0
		} else {
			panic("Number was not 0 or 1")
		}
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
