package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
	"github.com/sebastiangeiger/advent-of-code-2021/day_1"
	"github.com/sebastiangeiger/advent-of-code-2021/day_10"
	"github.com/sebastiangeiger/advent-of-code-2021/day_11"
	"github.com/sebastiangeiger/advent-of-code-2021/day_12"
	"github.com/sebastiangeiger/advent-of-code-2021/day_13"
	"github.com/sebastiangeiger/advent-of-code-2021/day_2"
	"github.com/sebastiangeiger/advent-of-code-2021/day_22"
	"github.com/sebastiangeiger/advent-of-code-2021/day_3"
	"github.com/sebastiangeiger/advent-of-code-2021/day_4"
	"github.com/sebastiangeiger/advent-of-code-2021/day_5"
	"github.com/sebastiangeiger/advent-of-code-2021/day_6"
	"github.com/sebastiangeiger/advent-of-code-2021/day_7"
	"github.com/sebastiangeiger/advent-of-code-2021/day_8"
	"github.com/sebastiangeiger/advent-of-code-2021/day_9"
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
	case 4:
		day_4.Run(problem)
	case 5:
		day_5.Run(problem)
	case 6:
		day_6.Run(problem)
	case 7:
		day_7.Run(problem)
	case 8:
		day_8.Run(problem)
	case 9:
		day_9.Run(problem)
	case 10:
		day_10.Run(problem)
	case 11:
		day_11.Run(problem)
	case 12:
		day_12.Run(problem)
	case 13:
		day_13.Run(problem)
	case 22:
		day_22.Run(problem)
	default:
		common.PrintNotImplemented(day, problem)
	}
}
