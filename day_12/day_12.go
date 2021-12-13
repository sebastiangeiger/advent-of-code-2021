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

func (node Node) CanBeRevisited() bool {
	firstLetter := node.name[0:1]
	return firstLetter == strings.ToUpper(firstLetter)
}

type Edge struct {
	start Node
	end   Node
}

func problem1() {
	fmt.Printf("Number of paths (test 1): %d\n", solveProblem1("day_12_test.input"))
	fmt.Printf("Number of paths (test 2): %d\n", solveProblem1("day_12_test_2.input"))
	fmt.Printf("Number of paths (test 3): %d\n", solveProblem1("day_12_test_3.input"))
	fmt.Printf("Number of paths (real): %d\n", solveProblem1("day_12.input"))
}

func problem2() {
	fmt.Printf("Number of paths (with revisit) (test 1): %d\n", solveProblem2("day_12_test.input"))
}

func solveProblem(path string, allowRevisit bool) int {
	edges := read(path)
	neighbors := makeNeighbors(edges)
	paths := [][]Node{[]Node{Node{"start"}}}
	i := 0
	for {
		i++
		paths = expand(paths, neighbors, allowRevisit)
		if allDone(paths) {
			break
		}
	}
	return len(paths)

}
func solveProblem1(path string) int {
	return solveProblem(path, false)
}

func solveProblem2(path string) int {
	return solveProblem(path, true)
}

func printPaths(paths [][]Node) {
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

func allDone(paths [][]Node) bool {
	done := true
	for _, path := range paths {
		last := path[len(path)-1]
		if last.name != "end" {
			done = false
			break
		}
	}
	return done
}

func expand(paths [][]Node, neighbors map[Node][]Node, allowRevisit bool) [][]Node {
	result := [][]Node{}
	for _, path := range paths {
		last := path[len(path)-1]
		if last.name == "end" {
			result = append(result, path)
		} else {
			for _, neighbor := range neighbors[last] {
				if neighbor.CanBeRevisited() {
					result = append(result, addNode(path, neighbor))
				} else if !includes(path, neighbor) {
					result = append(result, addNode(path, neighbor))
				}
			}
		}
	}
	return result
}

func addNode(path []Node, node Node) []Node {
	newPath := make([]Node, len(path)+1)
	copy(newPath, path)
	newPath[len(path)] = node
	return newPath
}

func includes(path []Node, node Node) bool {
	found := false
	for _, n := range path {
		if n.name == node.name {
			found = true
			break
		}
	}
	return found
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
