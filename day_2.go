package main

import (
	"fmt"
	"regexp"
	"strconv"
)

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
	testPosition := Position{0, 0}
	testPosition.ApplyNavigationSteps(readLinesFrom("day_2_test.input"))
	fmt.Printf("Multiplied (test): %d\n", testPosition.Multiply())
	realPosition := Position{0, 0}
	realPosition.ApplyNavigationSteps(readLinesFrom("day_2.input"))
	fmt.Printf("Multiplied (real): %d\n", realPosition.Multiply())
}

func day2problem2() {
	fmt.Printf("Implement me!\n")
}

type Position struct {
	horizontal int
	depth      int
}

type PositionDiff struct {
	horizontal int
	depth      int
}

func convertNavigationSteps(navigationSteps []string) []PositionDiff {
	positionDiffs := []PositionDiff{}
	regex := regexp.MustCompile("^(forward|down|up)\\s+(\\d+)$")
	for _, step := range navigationSteps {
		if regex.MatchString(step) {
			strings := regex.FindStringSubmatch(step)
			amount, err := strconv.Atoi(strings[2])
			direction := strings[1]
			if err != nil {
				panic(err)
			}
			positionDiffs = append(positionDiffs, makePositionDiff(direction, amount))
		} else {
			panic(fmt.Sprintf("Could not match regex to %#v", step))
		}
	}
	return positionDiffs
}

func makePositionDiff(direction string, amount int) PositionDiff {
	negativeAmount := -1 * amount
	switch direction {
	case "forward":
		return PositionDiff{horizontal: amount, depth: 0}
	case "up":
		return PositionDiff{horizontal: 0, depth: negativeAmount}
	case "down":
		return PositionDiff{horizontal: 0, depth: amount}
	default:
		panic(fmt.Sprintf("Don't know direction %s - amount %d\n", direction, amount))
	}
}

func (position *Position) ApplyNavigationSteps(navigationSteps []string) {
	positionDiffs := convertNavigationSteps(navigationSteps)
	for _, positionDiff := range positionDiffs {
		position.depth += positionDiff.depth
		position.horizontal += positionDiff.horizontal
	}
}

func (position Position) Multiply() int {
	return position.depth * position.horizontal
}
