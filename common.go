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
	lines := strings.Split(string(data), "\n")
	return lines
}
