package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadLinesFrom(path string) []string {
	data, fileError := os.ReadFile(filepath.Join("inputs", path))
	if fileError != nil {
		panic(fileError)
	}
	lines := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func PrintNotImplemented(day int, problem int) {
	fmt.Printf("Day %d - problem %d is not implemented yet", day, problem)
	os.Exit(1)
}
