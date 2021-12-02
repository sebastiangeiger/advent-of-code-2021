package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
	"github.com/sebastiangeiger/advent-of-code-2021/day_1"
	"github.com/sebastiangeiger/advent-of-code-2021/day_2"
	"github.com/sebastiangeiger/advent-of-code-2021/day_3"
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

func runDayProblem(day int, problem int) {
	switch day {
	case 1:
		day_1.Run(problem)
	case 2:
		day_2.Run(problem)
	case 3:
		day_3.Run(problem)
	default:
		common.PrintNotImplemented(day, problem)
	}
}
