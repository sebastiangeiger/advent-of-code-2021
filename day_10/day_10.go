package day_10

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
		common.PrintNotImplemented(7, problem)
	}
}

func problem1() {
	fmt.Printf("Illegals (test): %d\n", solveProblem1("day_10_test.input"))
}

func problem2() {
	fmt.Println("Day 10 - Problem 2")
}

func solveProblem1(path string) int {
	lines := common.ReadLinesFrom(path, false)
	sum := 0
	for _, line := range lines {
		sum += syntaxCheck(line)
	}
	return sum
}

type OpenClose int64

const (
	Open OpenClose = iota
	Close
)

func syntaxCheck(line string) int {
	// fmt.Printf("Syntax Check for '%s'\n", line)
	stack := []string{}
	for _, sym := range strings.Split(line, "") {
		openClose := openClose(sym)
		if openClose == Open {
			stack = append(stack, sym)
			// fmt.Printf("  '%s' was opening: '%s'\n", sym, strings.Join(stack, ""))
		} else {
			last := stack[len(stack)-1]
			remainder := stack[0 : len(stack)-1]
			if sym == matchingClose(last) {
				stack = remainder
				// fmt.Printf("  '%s' was closing: '%s' (deleted '%s')\n", sym, strings.Join(stack, ""), last)
			} else {
				// fmt.Printf("Found illegal '%s'\n", sym)
				return score(sym)
			}
		}
	}
	return 0
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

func score(symbol string) int {
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
