package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
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
	testPosition.ApplyNavigationSteps(common.ReadLinesFrom("day_2_test.input"))
	fmt.Printf("Multiplied (test): %d\n", testPosition.Multiply())
	realPosition := Position{0, 0}
	realPosition.ApplyNavigationSteps(common.ReadLinesFrom("day_2.input"))
	fmt.Printf("Multiplied (real): %d\n", realPosition.Multiply())
}

func day2problem2() {
	testPosition := PositionWithAim{0, 0, 0}
	testPosition.ApplyNavigationSteps(common.ReadLinesFrom("day_2_test.input"))
	fmt.Printf("Multiplied (test): %d\n", testPosition.Multiply())
	realPosition := PositionWithAim{0, 0, 0}
	realPosition.ApplyNavigationSteps(common.ReadLinesFrom("day_2.input"))
	fmt.Printf("Multiplied (real): %d\n", realPosition.Multiply())
}

type Position struct {
	horizontal int
	depth      int
}

type PositionDiff struct {
	horizontal int
	depth      int
}

func simpleConvertNavigationSteps(navigationSteps []string) []PositionDiff {
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
	positionDiffs := simpleConvertNavigationSteps(navigationSteps)
	for _, positionDiff := range positionDiffs {
		position.depth += positionDiff.depth
		position.horizontal += positionDiff.horizontal
	}
}

func (position Position) Multiply() int {
	return position.depth * position.horizontal
}

type PositionWithAim struct {
	horizontal int
	depth      int
	aim        int
}

func (position PositionWithAim) Multiply() int {
	return position.depth * position.horizontal
}

type OperationKind int64

const (
	AimAdjustment OperationKind = iota
	Movement
)

type Operation struct {
	kind   OperationKind
	amount int
}

func (position *PositionWithAim) ApplyNavigationSteps(navigationSteps []string) {
	operations := aimConvertNavigationSteps(navigationSteps)
	for _, operation := range operations {
		if operation.kind == AimAdjustment {
			position.aim += operation.amount
		} else if operation.kind == Movement {
			position.horizontal += operation.amount
			position.depth += position.aim * operation.amount
		} else {
			panic("Shouldn't get here!")
		}
	}
}

func aimConvertNavigationSteps(navigationSteps []string) []Operation {
	operations := []Operation{}
	regex := regexp.MustCompile("^(forward|down|up)\\s+(\\d+)$")
	for _, step := range navigationSteps {
		if regex.MatchString(step) {
			strings := regex.FindStringSubmatch(step)
			amount, err := strconv.Atoi(strings[2])
			direction := strings[1]
			if err != nil {
				panic(err)
			}
			operations = append(operations, makeOperation(direction, amount))
		} else {
			panic(fmt.Sprintf("Could not match regex to %#v", step))
		}
	}
	return operations
}

func makeOperation(direction string, amount int) Operation {
	negativeAmount := -1 * amount
	switch direction {
	case "forward":
		return Operation{kind: Movement, amount: amount}
	case "up":
		return Operation{kind: AimAdjustment, amount: negativeAmount}
	case "down":
		return Operation{kind: AimAdjustment, amount: amount}
	default:
		panic(fmt.Sprintf("Don't know direction %s - amount %d\n", direction, amount))
	}
}
