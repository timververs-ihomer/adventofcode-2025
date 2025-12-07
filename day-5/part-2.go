// Package main solves Advent of Code 2025 Day 5 puzzle part 2.
package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// Read the input.txt file and parse the product ranges
	inputData, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input.txt: %v", err)
	}

	lines := strings.Split(string(inputData), "\n")

	// loop through the input data and find the first line that is an empty string
	var indexOfEmptyLine int
	for index, line := range lines {
		if line == "" {
			indexOfEmptyLine = index
			break
		}
	}

	// everything before the empty line is the fresh product ranges
	freshProductRanges := lines[:indexOfEmptyLine]

}
