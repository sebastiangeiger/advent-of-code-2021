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
	fmt.Printf("%#v\n", readLinesFrom("day2_test.input"))
}

func day2problem2() {
	fmt.Printf("Implement me!\n")
}
