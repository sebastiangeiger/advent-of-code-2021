package common

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func ReadLinesFrom(path string, includeEmptyLines bool) []string {
	data, fileError := os.ReadFile(filepath.Join("inputs", path))
	if fileError != nil {
		panic(fileError)
	}
	allLines := strings.Split(string(data), "\n")
	if includeEmptyLines {
		return allLines
	} else {
		lines := []string{}
		for _, line := range allLines {
			if line != "" {
				lines = append(lines, line)
			}
		}
		return lines
	}
}

func PrintNotImplemented(day int, problem int) {
	fmt.Printf("Day %d - problem %d is not implemented yet", day, problem)
	os.Exit(1)
}

func InitializeArray(dx int, dy int) [][]int {
	result := make([][]int, dx)
	for i := 0; i < dx; i++ {
		result[i] = make([]int, dy)
	}
	return result
}

func InitializeBoolArray(dx int, dy int) [][]bool {
	result := make([][]bool, dx)
	for i := 0; i < dx; i++ {
		result[i] = make([]bool, dy)
	}
	return result
}

func Max(numbers ...int) int {
	max := math.MinInt
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
