package day_12

import (
	"fmt"
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
		common.PrintNotImplemented(12, problem)
	}
}

type Node struct {
	name string
}

type Edge struct {
	start Node
	end   Node
}

func problem1() {
	fmt.Printf("Number of paths (test): %d\n", solveProblem1("day_12_test.input"))
}

func problem2() {
	fmt.Println("Day 12 - Problem 1")
}

func solveProblem1(path string) int {
	edges := read(path)
	fmt.Printf("%#v\n", edges)
	return 1
}

func read(path string) []Edge {
	edges := []Edge{}
	for _, line := range common.ReadLinesFrom(path, false) {
		points := strings.Split(line, "-")
		edges = append(edges, Edge{Node{points[0]}, Node{points[1]}})
	}
	return edges
}
