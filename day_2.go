package main

import "fmt"

func runDay2(problem int) {
	switch problem {
	case 1:
		day2problem1()
	case 2:
		day2problem2()
	default:
		printNotImplemented(2, problem)
	}
}

func day2problem1() {
	navigationSteps := readLinesFrom("day2_test.input")
	fmt.Printf("%#v\n", navigationSteps)
	position := Position{0, 0}
	fmt.Printf("Start: %#v\n", position)
	position.ApplyNavigationSteps(navigationSteps)
	fmt.Printf("End: %#v\n", position)
}

func day2problem2() {
	fmt.Printf("Implement me!\n")
}

type Position struct {
	horizontal int
	depth      int
}

func (position *Position) ApplyNavigationSteps(navigationSteps []string) {
	position.depth = 1
	position.horizontal = 2
}
