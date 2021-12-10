package day_10

import (
	"fmt"
	"sort"
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
		common.PrintNotImplemented(7, problem)
	}
}

func problem1() {
	fmt.Printf("Illegals (test): %d\n", solveProblem1("day_10_test.input"))
	fmt.Printf("Illegals (real): %d\n", solveProblem1("day_10.input"))
}

func problem2() {
	fmt.Printf("Complete (test): %d\n", solveProblem2("day_10_test.input"))
	fmt.Printf("Complete (real): %d\n", solveProblem2("day_10.input"))
}

func solveProblem1(path string) int {
	lines := common.ReadLinesFrom(path, false)
	sum := 0
	for _, line := range lines {
		sum += syntaxCheck(line)
	}
	return sum
}

func solveProblem2(path string) int {
	lines := common.ReadLinesFrom(path, false)
	scores := []int{}
	for _, line := range lines {
		if syntaxCheck(line) == 0 {
			scores = append(scores, complete(line))
		}
	}
	sort.Ints(scores)
	return median(scores)
}

func median(input []int) int {
	if len(input)%2 == 1 {
		return input[len(input)/2]
	} else {
		panic("Need odd input")
	}
}

type OpenClose int64

const (
	Open OpenClose = iota
	Close
)

func syntaxCheck(line string) int {
	stack := []string{}
	for _, sym := range strings.Split(line, "") {
		openClose := openClose(sym)
		if openClose == Open {
			stack = append(stack, sym)
		} else {
			last := stack[len(stack)-1]
			remainder := stack[0 : len(stack)-1]
			if sym == matchingClose(last) {
				stack = remainder
			} else {
				return syntaxScore(sym)
			}
		}
	}
	return 0
}

func complete(line string) int {
	remainder := stillOpen(line)
	closed := make([]string, len(remainder))
	for i, _ := range remainder {
		closed[i] = matchingClose(remainder[len(remainder)-1-i])
	}
	score := 0
	for _, sym := range closed {
		score = score*5 + completeScore(sym)
	}
	return score
}

func stillOpen(line string) []string {
	stack := []string{}
	for _, sym := range strings.Split(line, "") {
		openClose := openClose(sym)
		if openClose == Open {
			stack = append(stack, sym)
		} else {
			last := stack[len(stack)-1]
			remainder := stack[0 : len(stack)-1]
			if sym == matchingClose(last) {
				stack = remainder
			} else {
				panic("This should not happen, please syntax check first!")
			}
		}
	}
	return stack
}

func openClose(symbol string) OpenClose {
	switch symbol {
	case "<", "{", "[", "(":
		return Open
	case ">", "}", "]", ")":
		return Close
	}
	panic(fmt.Sprintf("Don't know about '%s'", symbol))
}

func matchingClose(symbol string) string {
	switch symbol {
	case "<":
		return ">"
	case "{":
		return "}"
	case "[":
		return "]"
	case "(":
		return ")"
	}
	panic(fmt.Sprintf("Don't know about '%s'", symbol))
}

func syntaxScore(symbol string) int {
	switch symbol {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	panic(fmt.Sprintf("Don't know about '%s'", symbol))
}

func completeScore(symbol string) int {
	switch symbol {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	}
	panic(fmt.Sprintf("Don't know about '%s'", symbol))
}
