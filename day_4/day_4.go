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
	fmt.Printf("First win: UnmarkedSum x Draw (test): %d\n", solveProblem1("day_4_test.input"))
	fmt.Printf("First win: UnmarkedSum x Draw (real): %d\n", solveProblem1("day_4.input"))
}

func problem2() {
	fmt.Printf("Last win: UnmarkedSum x Draw (test): %d\n", solveProblem2("day_4_test.input"))
	fmt.Printf("Last win: UnmarkedSum x Draw (real): %d\n", solveProblem2("day_4.input"))
}

func solveProblem1(path string) int {
	lines := common.ReadLinesFrom(path, true)
	drawnNumbers, bingoBoards := makeDatastructures(lines)
	winningBoard, winningDraw := findFirstWin(bingoBoards, drawnNumbers)
	return winningDraw * winningBoard.UnmarkedSum()
}

func solveProblem2(path string) int {
	lines := common.ReadLinesFrom(path, true)
	drawnNumbers, bingoBoards := makeDatastructures(lines)
	winningBoard, winningDraw := findLastWin(bingoBoards, drawnNumbers)
	return winningDraw * winningBoard.UnmarkedSum()
}

func findFirstWin(bingoBoards []BingoBoard, drawnNumbers []int) (BingoBoard, int) {
	for _, drawnNumber := range drawnNumbers {
		for _, board := range bingoBoards {
			board.Mark(drawnNumber)
			if board.HasWon() {
				return board, drawnNumber
			}
		}
	}
	panic("Nothing has won and we're at the end of the drawn numbers")
}

func findLastWin(bingoBoards []BingoBoard, drawnNumbers []int) (BingoBoard, int) {
	for _, drawnNumber := range drawnNumbers {
		incompleteBoards := []BingoBoard{}
		for _, board := range bingoBoards {
			board.Mark(drawnNumber)
			if !board.HasWon() {
				incompleteBoards = append(incompleteBoards, board)
			}
		}
		if len(bingoBoards) == 1 && len(incompleteBoards) == 0 {
			return bingoBoards[0], drawnNumber
		} else {
			bingoBoards = incompleteBoards
		}
	}
	panic("We're at the end of the drawn numbers")
}

func makeDatastructures(lines []string) ([]int, []BingoBoard) {
	partitions := makePartitions(lines)
	drawnNumbers := makeDrawnNumbers(partitions[0][0])
	bingoBoards := makeBingoBoards(partitions[1:])
	return drawnNumbers, bingoBoards
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

func makeBingoBoards(partitions [][]string) []BingoBoard {
	boards := []BingoBoard{}
	for _, partition := range partitions {
		board := toIntMatrix(partition)
		marked := common.InitializeBoolArray(len(board), len(board[0]))
		bingoBoard := BingoBoard{board, marked}
		boards = append(boards, bingoBoard)
	}
	return boards
}

func toIntMatrix(lines []string) [][]int {
	dx := len(lines)
	dy := len(toIntLine(lines[0], " "))
	result := common.InitializeArray(dx, dy)
	for x := 0; x < dx; x++ {
		currentLine := toIntLine(lines[x], " ")
		if len(currentLine) != dy {
			panic(fmt.Sprintf("Expected lines[%d] to be %d long but was %d", x, dy, len(currentLine)))
		}
		result[x] = currentLine
	}
	return result
}

func toIntLine(line string, separator string) []int {
	result := []int{}
	for _, element := range strings.Split(line, separator) {
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
	return toIntLine(input, ",")
}
