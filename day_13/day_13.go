package day_13

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
		common.PrintNotImplemented(13, problem)
	}
}

type FoldInstruction struct {
	axis       string
	coordinate int
}

func problem1() {
	solveProblem1("day_13_test.input")
}

func solveProblem1(path string) {
	paper, foldInstructions := read(path)
	fmt.Printf("Paper: %#v\n", paper)
	fmt.Printf("FoldInstructions: %#v\n", foldInstructions)
}

func problem2() {
	fmt.Println("Day 13 - Problem 2")
}

func read(path string) ([][]int, []FoldInstruction) {
	lines := common.ReadLinesFrom(path, true)
	coordinates := [][]int{}
	foldInstructions := []FoldInstruction{}
	readCoordinates := true
	reg := regexp.MustCompile("^fold along (x|y)=(\\d+)$")
	for _, line := range lines {
		if readCoordinates && line != "" {
			coordinates = append(coordinates, common.ToIntLine(line, ","))
		} else if readCoordinates && line == "" {
			readCoordinates = false
		} else if !readCoordinates && line != "" {
			matches := reg.FindStringSubmatch(line)
			axis := matches[1]
			coordinate, err := strconv.Atoi(matches[2])
			if err == nil {
				foldInstructions = append(foldInstructions, FoldInstruction{axis, coordinate})
			} else {
				panic("Down here!")
			}
		}
	}
	return coordinates, foldInstructions
}
