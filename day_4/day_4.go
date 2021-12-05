package day_4

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
		common.PrintNotImplemented(4, problem)
	}
}

func problem1() {
	solveProblem1("day_4_test.input")
}

func problem2() {
	fmt.Println("Implement Day 4 - Problem 2")
}

func solveProblem1(path string) {
	lines := common.ReadLinesFrom(path, true)
	partitions := makePartitions(lines)
	drawnNumbers := makeDrawnNumbers(partitions[0][0])
	bingoBoards := makeBingoBoards(partitions[1:])
	// fmt.Printf("partitions: %#v", partitions)
	fmt.Printf("drawnNumbers: %#v", drawnNumbers)
	fmt.Printf("bingoBoards: %#v", bingoBoards)
}

func makePartitions(lines []string) [][]string {
	output := [][]string{}
	currentPartition := []string{}
	for _, line := range lines {
		if len(line) == 0 {
			output = append(output, currentPartition)
			currentPartition = []string{}
		} else {
			currentPartition = append(currentPartition, line)
		}
	}
	if len(currentPartition) > 0 {
		output = append(output, currentPartition)
	}
	return output
}

func makeBingoBoards(partitions [][]string) [][][]int {
	boards := [][][]int{}
	for _, partition := range partitions {
		board := toIntMatrix(partition)
		boards = append(boards, board)
	}
	return boards
}

func toIntMatrix(lines []string) [][]int {
	dx := len(lines)
	dy := len(toIntLine(lines[0]))
	result := common.InitializeArray(dx, dy)
	for x := 0; x < dx; x++ {
		currentLine := toIntLine(lines[x])
		if len(currentLine) != dy {
			panic(fmt.Sprintf("Expected lines[%d] to be %d long but was %d", x, dy, len(currentLine)))
		}
		result[x] = currentLine
	}
	return result
}

func toIntLine(line string) []int {
	result := []int{}
	for _, element := range strings.Split(line, " ") {
		if len(element) > 0 {
			number, err := strconv.Atoi(string(element))
			if err != nil {
				panic(err)
			} else {
				result = append(result, number)
			}
		}
	}
	return result
}

func makeDrawnNumbers(input string) []int {
	split := strings.Split(input, ",")
	output := make([]int, len(split))
	for i, s := range split {
		parsed, err := strconv.Atoi(s)
		if err == nil {
			output[i] = parsed
		} else {
			panic("Couldn't parse number")
		}
	}
	return output
}
