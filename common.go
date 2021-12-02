package main

import (
	"os"
	"strings"
)

func readLinesFrom(path string) []string {
	data, fileError := os.ReadFile(path)
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
