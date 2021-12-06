package day_5

import (
	"fmt"
	"regexp"
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
		common.PrintNotImplemented(5, problem)
	}
}

func problem1() {
	fmt.Printf("Problem1 (test): %d\n", solveProblem1("day_5_test.input"))
	fmt.Printf("Problem1 (real): %d\n", solveProblem1("day_5.input"))
}

func problem2() {
	fmt.Printf("Problem2 (test): %d\n", solveProblem2("day_5_test.input"))
	fmt.Printf("Problem2 (real): %d\n", solveProblem2("day_5.input"))
}

func solveProblem1(path string) int {
	strings := common.ReadLinesFrom(path, false)
	lines := onlyHorizontalAndVertical(makeLines(strings))
	field := makeField(lines)
	// printField(field)
	return findNumberOfOverlaps(field)
}

func solveProblem2(path string) int {
	strings := common.ReadLinesFrom(path, false)
	lines := onlyHorizontalVerticalAndDiagonal(makeLines(strings))
	field := makeField(lines)
	// printField(field)
	return findNumberOfOverlaps(field)
}

func makeField(lines []Line) [][]int {
	dx, dy := findDimensions(lines)
	field := common.InitializeArray(dx, dy)
	for _, line := range lines {
		points := line.ToPoints()
		for _, point := range points {
			field[point.x][point.y] += 1
		}
	}
	return field
}

func printField(field [][]int) {
	for _, line := range field {
		for _, cell := range line {
			if cell == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", cell)
			}
		}
		fmt.Printf("\n")
	}
}

func findNumberOfOverlaps(field [][]int) int {
	overlaps := 0
	for _, line := range field {
		for _, cell := range line {
			if cell >= 2 {
				overlaps += 1
			}
		}
	}
	return overlaps
}

func findDimensions(lines []Line) (int, int) {
	dx := 0
	dy := 0
	for _, line := range lines {
		dx = common.Max(dx, line.start.x, line.end.x)
		dy = common.Max(dy, line.start.y, line.end.y)
	}
	return dx + 1, dy + 1
}

func onlyHorizontalAndVertical(lines []Line) []Line {
	selectedLines := []Line{}
	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() {
			selectedLines = append(selectedLines, line)
		}
	}
	return selectedLines
}

func onlyHorizontalVerticalAndDiagonal(lines []Line) []Line {
	selectedLines := []Line{}
	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() || line.IsDiagonal() {
			selectedLines = append(selectedLines, line)
		}
	}
	return selectedLines
}

func makeLines(expressions []string) []Line {
	lines := []Line{}
	regex := regexp.MustCompile("^(\\d+),(\\d+)\\s+->\\s+(\\d+),(\\d+)$")
	for _, expression := range expressions {
		if regex.MatchString(expression) {
			matches := regex.FindStringSubmatch(expression)
			if len(matches) == 5 {
				startX := toIntOrPanic(matches[1])
				startY := toIntOrPanic(matches[2])
				endX := toIntOrPanic(matches[3])
				endY := toIntOrPanic(matches[4])
				line := Line{Point{startX, startY}, Point{endX, endY}}
				lines = append(lines, line)
			} else {
				panic("Expected 5 matches")
			}
		}
	}
	return lines
}

func toIntOrPanic(input string) int {
	number, err := strconv.Atoi(input)
	if err == nil {
		return number
	} else {
		panic("Couldn't convert number")
	}
}
