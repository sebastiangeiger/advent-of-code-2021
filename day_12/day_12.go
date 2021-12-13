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
	neighbors := makeNeighbors(edges)
	printNeighbors(neighbors)
	paths := [][]Node{[]Node{Node{"start"}}}
	printPaths(paths)
	return 1
}

func printPaths(paths [][]Node) {
	fmt.Printf("%d paths\n", len(paths))
	for i, path := range paths {
		fmt.Printf("%3d: ", i)
		nodeNames := make([]string, len(path))
		for j, node := range path {
			nodeNames[j] = node.name
		}
		fmt.Printf("%s\n", strings.Join(nodeNames, ","))
	}
}

func printNeighbors(neighbors map[Node][]Node) {
	for key, values := range neighbors {
		names := make([]string, len(values))
		for i, value := range values {
			names[i] = value.name
		}
		fmt.Printf("%s -> %s\n", key.name, strings.Join(names, ","))
	}
}

func makeNeighbors(edges []Edge) map[Node][]Node {
	neighbors := make(map[Node][]Node)
	for _, edge := range edges {
		addNeighbor(neighbors, edge.start, edge.end)
		addNeighbor(neighbors, edge.end, edge.start)
	}
	return neighbors
}

func addNeighbor(neighbors map[Node][]Node, key Node, value Node) {
	values, ok := neighbors[key]
	if ok {
		neighbors[key] = append(values, value)
	} else {
		neighbors[key] = []Node{value}
	}
}

func read(path string) []Edge {
	edges := []Edge{}
	for _, line := range common.ReadLinesFrom(path, false) {
		points := strings.Split(line, "-")
		edges = append(edges, Edge{Node{points[0]}, Node{points[1]}})
	}
	return edges
}
