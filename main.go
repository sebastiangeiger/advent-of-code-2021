package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		printHelp()
	} else {
		day, dayErr := strconv.Atoi(os.Args[1])
		problem, problemErr := strconv.Atoi(os.Args[2])
		if dayErr != nil {
			fmt.Println("Day was not an integer", dayErr)
			os.Exit(1)
		} else if problemErr != nil {
			fmt.Println("Problem was not an integer", dayErr)
			os.Exit(1)
		} else {
			runDayProblem(day, problem)
		}
	}

}

func printHelp() {
	fmt.Printf("Usage: '%s day problem'", os.Args[0])
	os.Exit(1)
}

func printNotImplemented(day int, problem int) {
	fmt.Printf("Day %d - problem %d is not implemented yet", day, problem)
	os.Exit(1)
}

func runDayProblem(day int, problem int) {
	switch day {
	case 1:
		runDay1(problem)
	case 2:
		runDay2(problem)
	default:
		printNotImplemented(day, problem)
	}
}
