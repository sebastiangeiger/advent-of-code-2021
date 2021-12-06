package day_5

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

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
}

func problem2() {
	fmt.Printf("Implement Day 5, Problem 2")
}

func solveProblem1(path string) int {
	strings := common.ReadLinesFrom(path, false)
	lines := makeLines(strings)
	fmt.Printf("%#v\n", lines)
	return 1
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
