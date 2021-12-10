package day_8

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/sebastiangeiger/advent-of-code-2021/common"
)

type Observation struct {
	signals []string
	output  []string
}

type DecodedObservation struct {
	signals []int
	output  []int
}

func (o Observation) Decode() DecodedObservation {
	possibleMappings := allMappings([]string{"a", "b", "c", "d", "e", "f", "g"})
	allDigits := append(o.signals, o.output...)

	viableMappings := []map[string]string{}
	for _, mapping := range possibleMappings {
		mappingIsViable := true
		for _, digit := range allDigits {
			applied := applyMapping(digit, mapping)
			_, err := signalToInt(applied)
			if err != nil {
				mappingIsViable = false
				break
			}
		}
		if mappingIsViable {
			viableMappings = append(viableMappings, mapping)
		}
	}
	if len(viableMappings) == 1 {
		viableMapping := viableMappings[0]
		decodedSignals := make([]int, len(o.signals))
		decodedOutput := make([]int, len(o.output))
		for i, signal := range o.signals {
			decoded, _ := signalToInt(applyMapping(signal, viableMapping))
			decodedSignals[i] = decoded
		}
		for i, output := range o.output {
			decoded, _ := signalToInt(applyMapping(output, viableMapping))
			decodedOutput[i] = decoded
		}
		return DecodedObservation{decodedSignals, decodedOutput}
	} else {
		panic(fmt.Sprintf("Expected to find 1 viable mapping but got %d", len(viableMappings)))
	}
}

func applyMapping(input string, mapping map[string]string) string {
	splitInput := strings.Split(input, "")
	result := make([]string, len(splitInput))
	for i, r := range splitInput {
		result[i] = mapping[r]
	}
	return strings.Join(result, "")
}

func allMappings(elements []string) []map[string]string {
	indexes := make([]int, len(elements))
	for i := range indexes {
		indexes[i] = i
	}
	permutations := permutations(indexes)
	result := []map[string]string{}
	for _, permutation := range permutations {
		currentMap := map[string]string{}
		for i, j := range permutation {
			key := elements[i]
			value := elements[j]
			currentMap[key] = value
		}
		result = append(result, currentMap)
	}
	return result
}

func Run(problem int) {
	switch problem {
	case 1:
		problem1()
	case 2:
		problem2()
	default:
		common.PrintNotImplemented(8, problem)
	}
}

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

func signalToInt(signal string) (int, error) {
	sortedSignal := sortInside(signal)
	switch sortedSignal {
	case "abcefg":
		return 0, nil
	case "cf":
		return 1, nil
	case "acdeg":
		return 2, nil
	case "acdfg":
		return 3, nil
	case "bcdf":
		return 4, nil
	case "abdfg":
		return 5, nil
	case "abdefg":
		return 6, nil
	case "acf":
		return 7, nil
	case "abcdefg":
		return 8, nil
	case "abcdfg":
		return 9, nil
	}
	return -1, errors.New(fmt.Sprintf("Could not match '%s' to an integer", sortedSignal))
}

func sortInside(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func interpretations(digit string) []int {
	switch len(digit) {
	case 2:
		return []int{1}
	case 3:
		return []int{7}
	case 4:
		return []int{4}
	case 5:
		return []int{2, 3, 5}
	case 6:
		return []int{0, 6, 9}
	case 7:
		return []int{8}
	}
	panic("Should not get here")
}

func problem1() {
	fmt.Printf("1,4,7,8 in output (test): %d\n", solveProblem1("day_8_test.input"))
	fmt.Printf("1,4,7,8 in output (real): %d\n", solveProblem1("day_8.input"))
}

func problem2() {
	fmt.Printf("Decoded (test): %d\n", solveProblem2("day_8_test.input"))
	fmt.Printf("Decoded (real): %d\n", solveProblem2("day_8.input"))
}

func solveProblem1(path string) int {
	observations := parseLines(common.ReadLinesFrom(path, false))
	sum := 0
	for _, observation := range observations {
		for _, digit := range observation.output {
			interpretations := interpretations(digit)
			if len(interpretations) == 1 {
				sum += 1
			}
		}
	}
	return sum
}

func toNumber(arr []int) int {
	number := 0.0
	for i, el := range arr {
		number += float64(el) * math.Pow(10, float64(len(arr)-i-1))
	}
	return int(number)
}

func solveProblem2(path string) int {
	observations := parseLines(common.ReadLinesFrom(path, false))
	sum := 0
	for _, observation := range observations {
		decoded := observation.Decode().output
		sum += toNumber(decoded)
	}
	return sum
}

func parseLines(lines []string) []Observation {
	observations := make([]Observation, len(lines))
	for i, line := range lines {
		modified := strings.Split(line, "|")
		if len(modified) == 2 {
			signalPatterns := parsePatterns(modified[0])
			output := parsePatterns(modified[1])
			observations[i] = Observation{signalPatterns, output}
		} else {
			panic("Expected 2")
		}
	}
	return observations
}

func parsePatterns(rawDigits string) []string {
	output := []string{}
	for _, digit := range strings.Split(rawDigits, " ") {
		if len(digit) > 0 {
			output = append(output, digit)
		}
	}
	return output
}
